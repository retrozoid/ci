package ci

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type tcTransport struct {
	headers map[string]string
}

var defSetting = &Setting{
	RetryMax:      10,
	RetryInterval: 500 * time.Millisecond,
}

// Setting ...
type Setting struct {
	RetryMax      int64
	RetryInterval time.Duration
}

// Client ...
type Client struct {
	url     string
	cli     *http.Client
	Setting *Setting
}

func (c *Client) req(method, path string, reader io.Reader) (*http.Response, error) {

	var err error
	var request *http.Request
	var response *http.Response

	var body []byte
	if reader != nil {
		if body, err = ioutil.ReadAll(reader); err != nil {
			return nil, err
		}
	}

	for retry := c.Setting.RetryMax; retry > 0; retry-- {
		if request, err = http.NewRequest(method, c.url+path, bytes.NewReader(body)); err != nil {
			return nil, err
		}
		response, err = c.cli.Do(request)
		if response == nil && err == nil {
			err = errors.New("nil response")
		}
		if err == nil {
			return response, nil
		}
		time.Sleep(c.Setting.RetryInterval)
	}

	return response, err
}

func (t *tcTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range t.headers {
		req.Header.Add(k, v)
	}
	return http.DefaultTransport.RoundTrip(req)
}

// New ...
func New(host, user, password string) *Client {
	tt := &tcTransport{
		headers: map[string]string{
			"Content-Type":  "application/json",
			"Accept":        "application/json",
			"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+password)),
		},
	}
	return &Client{
		url:     fmt.Sprintf("%s/httpAuth/app/rest/", host),
		cli:     &http.Client{Transport: tt},
		Setting: defSetting,
	}
}

func (c *Client) get(path string, response interface{}) error {
	return c.request(http.MethodGet, path, nil, response)
}

func (c *Client) post(path string, body interface{}, response interface{}) error {
	return c.request(http.MethodPost, path, body, response)
}

func (c *Client) request(method, path string, body interface{}, response interface{}) error {
	var err error
	var b []byte

	if body != nil {
		b, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}

	r, err := c.req(method, path, bytes.NewReader(b))
	if err != nil {
		return err
	}

	res := new(bytes.Buffer)
	res.ReadFrom(r.Body)
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return errors.New(res.String())
	}
	if response == nil {
		return nil
	}
	return json.NewDecoder(bytes.NewReader(res.Bytes())).Decode(response)
}

package ci

import "fmt"

var (
	defBranchName = "<default>"
)

// NewProperty ...
func NewProperty(name, value string) *Property {
	return &Property{Name: name, Value: value}
}

// ConfQuery ...
type ConfQuery struct {
	build *Build
	c     *Client
}

// BuildType select buildConfiguration
func (c *Client) BuildType(buildTypeID string) *ConfQuery {
	return &ConfQuery{c: c, build: &Build{BuildTypeID: buildTypeID}}
}

// WithProperties ...
func (q *ConfQuery) WithProperties(vars ...*Property) *ConfQuery {
	if q.build.Properties == nil {
		q.build.Properties = &Properties{}
	}
	q.build.Properties.Property = append(q.build.Properties.Property, vars...)
	return q
}

// Queue ...
func (q *ConfQuery) Queue() (*Build, error) {
	resp := new(Build)
	return resp, q.c.post("buildQueue", q.build, resp)
}

// Property get buildConfiguration property by name
func (q *ConfQuery) Property(name string) (*Property, error) {
	property := new(Property)
	return property, q.c.get(fmt.Sprintf("buildTypes/id:%s/parameters/%s", q.build.BuildTypeID, name), property)
}

// SetProperty set buildConfiguration property
func (q *ConfQuery) SetProperty(p *Property) error {
	return q.c.post(fmt.Sprintf("buildTypes/id:%s/parameters", q.build.BuildTypeID), p, nil)
}

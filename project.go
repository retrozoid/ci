package ci

import "fmt"

// ProjectQuery ...
type ProjectQuery struct {
	projectID string
	c         *Client
}

// Project select project
func (c *Client) Project(projectID string) *ProjectQuery {
	return &ProjectQuery{c: c, projectID: projectID}
}

// Property get project property by name
func (q *ProjectQuery) Property(name string) (*Property, error) {
	property := new(Property)
	return property, q.c.get(fmt.Sprintf("projects/%s/parameters/%s", q.projectID, name), property)
}

// SetProperty set project property
func (q *ProjectQuery) SetProperty(p *Property) error {
	return q.c.post(fmt.Sprintf("projects/%s/parameters", q.projectID), p, nil)
}

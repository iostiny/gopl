// Pckage github provides a Go API for the Github issue tracker.
package main

import (
	"fmt"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"
const APIURL = "https://api.github.com"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number       int
	HTMLURL      string `json:"html_url"`
	Title, State string
	User         *User
	CreateAt     time.Time `json:"created_at"`
	Body         string
}

func (i Issue) CacheURL() string {
	return fmt.Sprintf("/issues/%d", i.Number)
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues querries the Github issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	//!-
	q := url.QueryEscape(strings.Join(terms, " "))
	//q is :repo%3Agolang%2Fgo+is%3Aopen+json+decoder
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	//!+

	// req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	// if err != nil {
	// 	return nil, err
	// }

	// req.Header.Set(
	// 	"Accept", "application/vnd.github.v3.text-match+json")
	// resp, err := http.DefaultClient.Do(req)

	// We must close resp.Body on all execution paths.
	// (Chapter 5 present 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

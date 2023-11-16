// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package github

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	//!-
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	//
	//   req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	//   if err != nil {
	//       return nil, err
	//   }
	//   req.Header.Set(
	//       "Accept", "application/vnd.github.v3.text-match+json")
	//   resp, err := http.DefaultClient.Do(req)
	//!+

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
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

// Paginate through the results.
func PaginateSearchIssues(result *IssuesSearchResult) (*Issue, error) {
	reader := bufio.NewReader(os.Stdin)
	page := 5
	for count := 0; count < result.TotalCount; count += page {
		// Don't go out of bounds of the slice.
		if remaining := result.TotalCount - count; page > remaining {
			page = remaining
		}
		for _, issue := range result.Items[count : count+page] {
			fmt.Printf("#%-5d %9.9s %.55s\r\n", issue.Number, issue.User.Login, issue.Title)
		}
		if result.TotalCount-count > page {
			fmt.Printf("%d of %d \r\n Select an issue or press n to continue", count+page, result.TotalCount)
			input, error := reader.ReadString('\n')

			if error != nil {
				return nil, error
			}

			// @TODO Fix this to use terminal buffer now.
			if input == "n" {
				continue
			}
			if isNumber(input) {

			}
		}
	}
	return nil, nil
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

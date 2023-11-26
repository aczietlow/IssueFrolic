// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package github

import (
	"aczietlow/IssueFrolic/net"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
	ItemsMap   map[int]*Issue
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
	request, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	if err != nil {
		return nil, err
	}

	resp, err := net.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var SearchResult IssuesSearchResult

	githubJsonData, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error with GET Request: %v\n", err)
	}

	if err := json.Unmarshal(githubJsonData, &SearchResult); err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
	}

	issueMap := make(map[int]*Issue)
	for _, issue := range SearchResult.Items {
		issueMap[issue.Number] = issue
	}

	SearchResult.ItemsMap = issueMap

	return &SearchResult, nil
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

		}
	}
	return nil, nil
}

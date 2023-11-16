package main

import (
	"aczietlow/IssueFrolic/github"
	"fmt"
	"log"
	"os"
)

func main() {

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d Issues: \n", result.TotalCount)

	github.PaginateSearchIssues(result)

	// @TODO capture a specific issue to work with?
}

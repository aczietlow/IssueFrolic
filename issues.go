package main

import (
	"aczietlow/IssueFrolic/cli"
	"aczietlow/IssueFrolic/github"
	"fmt"
	"log"
	"os"
)

func main() {

	terminal := cli.NewTerminal()
	//options := terminal.Args.Options
	defer terminal.Restore()

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d Issues: \r\n", result.TotalCount)

	github.PaginateSearchIssues(result)

	// @TODO capture a specific issue to work with?
}

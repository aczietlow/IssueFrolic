package main

import (
	"aczietlow/IssueFrolic/cli"
	"aczietlow/IssueFrolic/config"
	"aczietlow/IssueFrolic/github"
	"aczietlow/IssueFrolic/net"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	config, err := config.LoadConfig("config.json")

	terminal := cli.NewTerminal()
	net.NewClient(config.Token)

	//options := terminal.Args.Options
	defer terminal.Restore()

	terminal.Terminal.Write([]byte(config.Username + "\r\n"))

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d Issues: \r\n", result.TotalCount)

	github.PaginateSearchIssues(result)

	input, _ := terminal.Terminal.ReadLine()

	terminal.Terminal.Write([]byte(input + "\r\n"))
	// @TODO capture a specific issue to work with?

	if isInt, key := IsNumber(input); isInt {
		issue, exists := result.ItemsMap[key]
		if exists {
			terminal.Terminal.Write([]byte(issue.Body + "\r\n"))
			terminal.Terminal.Write([]byte("Click here for more information:" + issue.HTMLURL + "\r\n"))
		}
	}
}

func IsNumber(s string) (bool, int) {
	num, err := strconv.Atoi(s)
	isInt := err == nil
	return isInt, num
}

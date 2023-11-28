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

	conf, err := config.LoadConfig("config.json")

	terminal := cli.NewTerminal()
	net.NewClient(conf.Token)

	//options := terminal.Args.Options
	defer terminal.Restore()

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d Issues: \r\n", result.TotalCount)

	github.PaginateSearchIssues(result)

	// This is insane, we need some sort of auto discovery pattern here, we can't write a giant switch case to rule them all.
	for {
		input, _ := terminal.Terminal.ReadLine()
		if isInt, key := IsNumber(input); isInt {
			issue, exists := result.ItemsMap[key]
			if exists {
				terminal.Terminal.Write([]byte(issue.Body + "\r\n"))
				terminal.Terminal.Write([]byte("Click here for more information:" + issue.HTMLURL + "\r\n"))
			} else {
				terminal.Terminal.Write([]byte("Provided issue not in returned list, try again or enter a new search \r\n"))
			}
		}

		if input == `exit` {
			break
		}
	}
}

func IsNumber(s string) (bool, int) {
	num, err := strconv.Atoi(s)
	isInt := err == nil
	return isInt, num
}

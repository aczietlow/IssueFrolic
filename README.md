# Issue Frolic

Tiptoe through the tulips that are gitHub issues.

A tool to let users explore, search, create, read, update, and close GitHub issues from the command line. 

## Why

Shits and Giggles, per normal

Heavily inspired and borrowed from an exercise from "The GO Programming Language" by Alan A. A. Donovan and Brian W. Kernighan

## Getting Started

Create a config.json using the contents of config.default.json as a template. 

Read permissions are permissible with out an access token, however write permissions will require token access. To create a token perform the following steps:

1. Create a Personal Access Token
- Go to your GitHub account settings.
- Navigate to "Developer settings" > "Personal access tokens".
- Click on "Generate new token".
- Select the scopes or permissions you'd like the token to have and generate the token.
- Copy the generated token. Make sure to save it securely as it's shown only once.
2. Using the Token in Your Code:
- Add your token to the config.json file.

## Usages

Very much WIP

`go run issues.go $githubSearchSyntax`

Search for all issues on a given repo
`go run issues.go repo:aczietlow/goGoPing`

Search for all issues on a given repo
`go run issues.go repo:aczietnlow/goGoPing`

Search for issues by a given user
`go run issues.go user:aczietlow`


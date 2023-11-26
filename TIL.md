# TIL

## Go Experimental packages

golang.org/x/<foo> are experimental packages, they can introduce breaking changes or go away entirely. 

## Terminal Apps in Go

Place a terminal in raw mode removes any kind of built in text formatting. For example "\n" will get interrupted literally as a new line, despite the expected behavior today being to enter a new line and carriage return to the beginning of hte terminal limits.

"What is a VT100" terminal

Video terminal from the 1970's, popularize b/c it supported ANSI standard code standards. Fun fact this is also where we get the 80 character limitation from as the terminal was 80 column wide and 24 lines in height.



## Searching on Github

[source](https://docs.github.com/en/search-github/getting-started-with-searching-on-github/understanding-the-search-syntax)

The search api supports the following syntax:

`QUALIFIER:USERNAME`


Qualifier	Example
type:pr	cat type:pr matches pull requests with the word "cat."
type:issue	github commenter:defunkt type:issue matches issues that contain the word "github," and have a comment by @defunkt.
is:pr	event is:pr matches pull requests with the word "event."
is:issue	is:issue label:bug is:closed matches closed issues with the label "bug."

Qualifier	Example
in:title	warning in:title matches issues with "warning" in their title.
in:body	error in:title,body matches issues with "error" in their title or body.
in:comments	shipit in:comments matches issues mentioning "shipit" in their comments.


Qualifier	Example
user:USERNAME	user:defunkt ubuntu matches issues with the word "ubuntu" from repositories owned by @defunkt.
org:ORGNAME	org:github matches issues in repositories owned by the GitHub organization.
repo:USERNAME/REPOSITORY	repo:mozilla/shumway created:<2012-03-01 matches issues from @mozilla's shumway project that were created before March 2012.

[more](https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests)
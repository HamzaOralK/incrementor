package main

import (
	"flag"
	"fmt"
	"incrementor/clients"
	"incrementor/utility"
	"log"
	"strings"
)

func main() {
	// input will be the release branch name as the following <prefix><separator>d.d.x
	// d.d.x will be the semver but the patch will be the x since we are going to increment
	// each patch with the next one.

	owner := flag.String("owner", "", "Owner organisation or person of the branch.")
	repository := flag.String("repository", "", "Repository name.")
	prefix := flag.String("prefix", "", "prefix used in release branch name.")
	separator := flag.String("separator", "", "Separator used in release branch name.")
	branch := flag.String("branch", "", "Release branch name.")
	flag.Parse()

	normalizedBranch := strings.Replace(*branch, fmt.Sprintf("%s%s", *prefix, *separator), "", 1)
	// First validate the branch as a valid input
	matched := utility.Validate(normalizedBranch, "\\d+\\.\\d+\\.x$")
	if matched == false {
		log.Fatalln("Branch is not matching with the following syntax of \"number.number.x\"")
	}
	// Get the tags and filter it by using the branch name return incremented one
	githubClient := clients.CreateGithubClient()
	tagToBeCreated := githubClient.GetIncrementedTag(owner, repository, normalizedBranch)
	fmt.Println(tagToBeCreated)
}

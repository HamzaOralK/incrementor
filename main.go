package main

import (
	"flag"
	"fmt"
	"incrementor/clients"
	"incrementor/utility"
	"log"
)

func main() {
	// input will be the release branch name as the following <prefix><separator>d.d.x
	// d.d.x will be the semver but the patch will be the x since we are going to increment
	// each patch with the next one.

	owner := flag.String("owner", "", "Owner organisation or person of the branch.")
	repository := flag.String("repository", "", "Repository name.")
	prefix := flag.String("prefix", "", "prefix used in release branch name.")
	separator := flag.String("separator", "", "Separator used in release branch name.")
	branchName := flag.String("branch", "", "Release branch name.")
	flag.Parse()

	// removes prefix and seperator from the branch name
	normalizedBranchName := utility.NormalizeBranchName(*branchName, *prefix, *separator)

	// First validate the branch as a valid input
	matched := utility.Validate(normalizedBranchName, "\\d+\\.\\d+\\.x$")
	if matched == false {
		log.Fatalln("Branch is not matching with the following syntax of \"number.number.x\"")
	}
	// Get the tags and filter it by using the branch name return incremented one
	githubClient := clients.CreateGithubClient()
	tagToBeCreated := githubClient.GetIncrementedTag(owner, repository, normalizedBranchName)
	fmt.Println(tagToBeCreated)
	test
}

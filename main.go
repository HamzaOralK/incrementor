package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/go-github/v43/github"
	semver2 "github.com/rogpeppe/go-internal/semver"
	"golang.org/x/oauth2"
	"log"
	"os"
	"regexp"
	"strconv"
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
	matched := validate(normalizedBranch, "\\d+\\.\\d+\\.x$")
	if matched == false {
		log.Fatalln("Branch is not matching with the following syntax of \"number.number.x\"")
	}
	// Get the tags and filter it by using the branch name return incremented one
	ctx := context.Background()
	client := createGithubClient(ctx)
	tagToBeCreated := getIncrementedTag(ctx, client, owner, repository, normalizedBranch)
	fmt.Println(tagToBeCreated)
}

func validate(input string, regex string) bool {
	matched, err := regexp.MatchString(regex, input)
	if err != nil {
		log.Fatalln(err)
	}
	return matched
}

func createGithubClient(ctx context.Context) *github.Client {

	token, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		log.Fatalln("GITHUB_TOKEN is not set")
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func getIncrementedTag(ctx context.Context, client *github.Client, owner *string, repository *string, branch string) string {
	nextPage := 1
	lastTag := ""
	for {
		tags, resp, _ := client.Repositories.ListTags(ctx, *owner, *repository, &github.ListOptions{Page: nextPage, PerPage: 100})
		for _, tag := range tags {
			regex := strings.Replace(branch, "x", "\\d+$", 1)
			matched := validate(*tag.Name, regex)
			if matched {
				if lastTag == "" {
					lastTag = *tag.Name
				} else {
					lastTag = getMaxSemver(lastTag, *tag.Name)
				}
			}
		}
		if resp.NextPage == 0 {
			break
		} else {
			nextPage = resp.NextPage
		}
	}
	if lastTag == "" {
		return strings.Replace(branch, "x", "0", 1)
	} else {
		return incrementPatch(lastTag)
	}
}

func getMaxSemver(v string, w string) string {
	return strings.Replace(semver2.Max(fmt.Sprintf("v%s", v), fmt.Sprintf("v%s", w)), "v", "", 1)
}

func incrementPatch(tag string) string {
	semverArray := strings.Split(tag, ".")
	i, _ := strconv.Atoi(semverArray[len(semverArray)-1])
	i = i + 1
	semverArray[len(semverArray)-1] = strconv.Itoa(i)
	return strings.Join(semverArray, ".")
}

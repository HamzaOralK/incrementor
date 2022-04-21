package clients

import (
	"context"
	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
	"incrementor/utility"
	"log"
	"os"
	"strings"
)

type GitHubClient struct {
	client *github.Client
	ctx    context.Context
}

func CreateGithubClient() GitHubClient {
	token, ok := os.LookupEnv("GITHUB_TOKEN")
	ctx := context.TODO()
	if !ok {
		log.Fatalln("GITHUB_TOKEN is not set")
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return GitHubClient{
		client: github.NewClient(tc),
		ctx:    ctx,
	}
}

func (g *GitHubClient) GetIncrementedTag(owner *string, repository *string, branch string) string {
	nextPage := 1
	lastTag := ""
	for {
		tags, resp, _ := g.client.Repositories.ListTags(g.ctx, *owner, *repository, &github.ListOptions{Page: nextPage, PerPage: 100})
		for _, tag := range tags {
			regex := strings.Replace(branch, "x", "\\d+$", 1)
			matched := utility.Validate(*tag.Name, regex)
			if matched {
				if lastTag == "" {
					lastTag = *tag.Name
				} else {
					lastTag = utility.GetMaxSemver(lastTag, *tag.Name)
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
		return utility.IncrementPatch(lastTag)
	}
}

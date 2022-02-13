package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type RepoDetails struct {
	name        string
	language    string
	url         string
	description string
	lastUpdated string
	size        int
}

func (rd RepoDetails) String() string {
	return fmt.Sprintf(rd.name)
}

func getRepos(page int, token string, organisation string) ([]*github.Repository, error) {

	const PER_PAGE = 100
	// list all repositories for the authenticated user
	var options *github.RepositoryListByOrgOptions = new(github.RepositoryListByOrgOptions)
	options.PerPage = PER_PAGE
	options.Page = page

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	repos, _, err := client.Repositories.ListByOrg(ctx, organisation, options)
	if err != nil {
		return nil, err
	}
	return repos, nil
}

func extractRepoDetails(repo *github.Repository) RepoDetails {
	var details RepoDetails
	details.name = repo.GetName()
	details.language = repo.GetLanguage()
	year, month, day := repo.GetPushedAt().Date()
	m := strconv.Itoa(int(month))
	d := strconv.Itoa(day)
	y := strconv.Itoa(year)
	details.lastUpdated = d + "/" + m + "/" + y
	details.description = repo.GetDescription()
	details.url = repo.GetCloneURL()
	details.size = repo.GetSize()
	return details
}

func processRepoDetails(repo RepoDetails) []string {
	var details []string
	details = append(details, repo.name)
	details = append(details, repo.language)
	details = append(details, repo.description)
	details = append(details, repo.url)
	details = append(details, repo.lastUpdated)
	details = append(details, strconv.Itoa(repo.size))
	return details
}

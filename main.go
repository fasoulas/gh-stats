package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var currentPage = 1
	const maxNoOfPages = 10

	tokenPtr := flag.String("token", "", "Github personal access token")
	token := os.Getenv("GH_TOKEN")
	orgPtr := flag.String("org", "", "Organisation to filter")
	flag.Parse()

	if token == "" && *tokenPtr == "" {
		fmt.Println("You must supply a valid Github access token via argument -token=")
		fmt.Println("Or you must set and export an env variable GH_TOKEN to use")
		os.Exit(1)
	}

	if *orgPtr == "" {
		fmt.Println("You must supply an organisation to filter by via argument -org=")
		os.Exit(2)
	}

	//pick one of the set tokens to use, prioritising command line arg
	var ghToken string
	if *tokenPtr != "" {
		ghToken = *tokenPtr
	} else {
		ghToken = token
	}

	var githubRepos []RepoDetails

	for {
		repos, err := getRepos(currentPage, ghToken, *orgPtr)
		if err != nil {
			fmt.Println("Error from Github. Please check your token and organisation you have used")
			fmt.Println(err)
			os.Exit(3)
		}

		if len(repos) == 0 {
			break
		}

		if currentPage == maxNoOfPages {
			break
		}

		for i := range repos {
			repoDetails := extractRepoDetails(repos[i])
			githubRepos = append(githubRepos, repoDetails)
		}
		currentPage++
	}

	fmt.Println(getCVSHeader("name", "lang", "description,", "url", "lastUpdated", "size"))
	for i := range githubRepos {
		fmt.Println(getCVSLine(processRepoDetails(githubRepos[i])))
	}
}

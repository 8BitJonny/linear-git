package main

import (
	"errors"
	"github.com/8bitjonny/linear-git/auth"
	"github.com/8bitjonny/linear-git/config"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
)

func main() {
	app := &cli.App{
		Name:    "gli",
		Version: "v0.1.0",
		Usage:   "A nice git command line helper for linear",
		Commands: []*cli.Command{
			{
				Name:        "login",
				Description: "Authenticate gli to read issue information from linear",
				Usage:       "login to linear",
				Action: func(cCtx *cli.Context) error {
					server := auth.CreateAuthCallbackServer()
					if err := auth.OpenAuthScreen(); err != nil {
						return err
					}

					println("Complete authorization in opened browser tab\nWaiting...")
					token := server.GetAuthToken()

					// Save token
					appConfig, err := config.ReadFromFilesystem()
					if err != nil {
						return err
					}
					appConfig.AuthToken = token
					err = appConfig.WriteToFilesystem()
					if err != nil {
						return err
					}
					println("Logged in successfully")
					return nil
				},
			}, {
				Name:        "branch",
				Description: "Create branch for specified linear ticket. It just wraps 'git checkout -b' meaning it'll create a new branch from your current one.",
				Usage:       "create branch for linear ticket",
				ArgsUsage:   "[linearIssueId]",
				Action: func(cCtx *cli.Context) error {
					issueID := cCtx.Args().First()

					appConfig, err := config.ReadFromFilesystem()
					if err != nil {
						return err
					} else if appConfig.AuthToken == "" {
						return errors.New("not authorized yet. You have to run 'gli login' first")
					}

					branchName, err := GetBranchName(issueID, appConfig.AuthToken)
					if err != nil {
						return err
					}

					if err = exec.Command("git", "checkout", "-b", branchName).Start(); err != nil {
						return err
					}
					println("Switched to Branch: ", branchName)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

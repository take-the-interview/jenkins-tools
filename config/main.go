package config

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	Ctx    context.Context
	Client *github.Client
)

func Auth(authToken string) {
	Ctx = context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: authToken},
	)
	tc := oauth2.NewClient(Ctx, ts)

	Client = github.NewClient(tc)
}

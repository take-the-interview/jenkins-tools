package config

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/bndr/gojenkins"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"net/http"
	"os"
)

var (
	Ctx     context.Context
	Client  *github.Client
	Jenkins *gojenkins.Jenkins
)

func Auth(authToken string) {
	Ctx = context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: authToken},
	)
	tc := oauth2.NewClient(Ctx, ts)

	Client = github.NewClient(tc)
}

func JenkinsAuth(serverUrl, user, pass string) {
	// fmt.Printf("%s %s %s\n", serverUrl, user, pass)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: tr}

	var err error
	Jenkins, err = gojenkins.CreateJenkins(httpClient, serverUrl, user, pass).Init()

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		return
	}
}

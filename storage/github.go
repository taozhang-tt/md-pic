package storage

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	"github.com/taozhang-tt/mdpic/config"
)

type Github struct {
	Token string
	Owner string
	Repo  string
}

func NewGithub(config *config.Config) Storage {
	return &Github{
		Token: config.GithubToken,
		Repo:  config.GithubRepo,
		Owner: config.GithubOwner,
	}
}

func (g *Github) UploadBytes(path string, bs []byte) (string, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: g.Token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	opts := &github.RepositoryContentFileOptions{
		Message: github.String("文件上传"),
		Content: bs,
		//Committer: &github.CommitAuthor{Name: github.String("TT"), Email: github.String("taozhang.tt@gmail.com")},
	}
	_, _, err := client.Repositories.CreateFile(ctx, g.Owner, g.Repo, path, opts)
	if err != nil {
		return "", err
	}
	return "https://raw.githubusercontent.com/" + g.Owner + "/" + g.Repo + "/main", nil
}

func (g *Github) DeleteObject(key string) error {
	return nil
}

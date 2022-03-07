package storage

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-github/github"
	. "github.com/smartystreets/goconvey/convey"
	"golang.design/x/clipboard"
	"golang.org/x/oauth2"
)

func Test(t *testing.T) {
	Convey("Test github upload", t, func() {
		Upload()
	})
}

func Upload() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	bs := clipboard.Read(clipboard.FmtImage)

	opts := &github.RepositoryContentFileOptions{
		Message: github.String("测试传文件"),
		Content: bs,
		//Committer: &github.CommitAuthor{Name: github.String("TT"), Email: github.String("taozhang.tt@gmail.com")},
	}
	r1, r2, err := client.Repositories.CreateFile(ctx, "taozhang-tt", "pictures", "1646631812.png", opts)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("r1: %+v\n, r2: %+v\n", r1, r2)

}

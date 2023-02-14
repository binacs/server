package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/github"

	"github.com/binacs/server/config"
	"github.com/binacs/server/types"
	"github.com/binacsgo/log"
	"github.com/binacsgo/pastebin"
)

// BlogServiceImpl Web PasteBin service implement
type BlogServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"WebLogger"`

	instance pastebin.PasteBin
	prefix   string

	gitOwner string
	gitRepo  string
	opt      github.RepositoryContentGetOptions
	client   *github.Client
}

// AfterInject do inject
func (bs *BlogServiceImpl) AfterInject() error {
	bs.instance = new(pastebin.MarkDownImpl)
	bs.prefix = bs.Config.WebConfig.GetDomain() + "/blog/"
	bs.gitOwner, bs.gitRepo = "binacs", "blog"
	bs.opt = github.RepositoryContentGetOptions{Ref: "main"}
	bs.client = github.NewClient(nil)
	return nil
}

func (bs *BlogServiceImpl) URLSearch(uri string) (types.Blog, error) {
	fileContent, _, _, _ := bs.client.Repositories.GetContents(context.Background(), bs.gitOwner, bs.gitRepo, uri, &bs.opt)
	if fileContent == nil {
		return types.Blog{}, fmt.Errorf("file not found")
	}
	context, _ := fileContent.GetContent()
	return types.Blog{
		Name:    fileContent.GetName(),
		Url:     fileContent.GetHTMLURL(),
		Content: context,
	}, nil
}

func (bs *BlogServiceImpl) RecentBlogs() ([]types.Blog, error) {
	_, directoryContent, _, _ := bs.client.Repositories.GetContents(context.Background(), bs.gitOwner, bs.gitRepo, "", &bs.opt)
	blogs := []types.Blog{}
	for i := range directoryContent {
		file := directoryContent[i]
		url := file.GetDownloadURL()
		if strings.HasSuffix(url, ".md") {
			blogs = append(blogs, types.Blog{
				Name: file.GetName(),
				Url:  bs.prefix + file.GetName(),
			})
		}
	}
	return blogs, nil
}

func (ps *BlogServiceImpl) Parse(content string) string {
	res, err := ps.instance.ParseContent(content)
	if err != nil {
		return err.Error()
	}
	return res
}

package service

import (
	"reflect"
	"testing"

	"github.com/binacs/server/config"
	"github.com/binacs/server/types"
	"github.com/binacsgo/log"
	"github.com/binacsgo/pastebin"
	"github.com/google/go-github/github"
)

func getNormalConfigWithBlogs() *config.Config {
	return &config.Config{
		WebConfig: config.WebConfig{
			SSLRedirect: false,
			Host:        "test.binacs.space",
		},
	}
}

func TestBlogServiceImpl_RecentBlogs(t *testing.T) {
	type fields struct {
		Config   *config.Config
		Logger   log.Logger
		instance pastebin.PasteBin
		prefix   string
		gitOwner string
		gitRepo  string
		opt      github.RepositoryContentGetOptions
		client   *github.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    []types.Blog
		wantErr bool
	}{
		{
			name: "RecentBlogs",
			fields: fields{
				Config:   getNormalConfigWithBlogs(),
				Logger:   log.NewNopLogger(),
				instance: new(pastebin.MarkDownImpl),
				prefix:   "test.binacs.space" + "/blog/",
				gitOwner: "binacs",
				gitRepo:  "blog",
				opt:      github.RepositoryContentGetOptions{Ref: "main"},
				client:   github.NewClient(nil),
			},
			want:    []types.Blog{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BlogServiceImpl{
				Config:   tt.fields.Config,
				Logger:   tt.fields.Logger,
				instance: tt.fields.instance,
				prefix:   tt.fields.prefix,
				gitOwner: tt.fields.gitOwner,
				gitRepo:  tt.fields.gitRepo,
				opt:      tt.fields.opt,
				client:   tt.fields.client,
			}
			got, err := bs.RecentBlogs()
			if (err != nil) != tt.wantErr {
				t.Errorf("BlogServiceImpl.RecentBlogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("BlogServiceImpl.RecentBlogs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlogServiceImpl_URLSearch(t *testing.T) {
	type fields struct {
		Config   *config.Config
		Logger   log.Logger
		instance pastebin.PasteBin
		prefix   string
		gitOwner string
		gitRepo  string
		opt      github.RepositoryContentGetOptions
		client   *github.Client
	}
	type args struct {
		uri string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    types.Blog
		wantErr bool
	}{
		{
			name: "RecentBlogs",
			fields: fields{
				Config:   getNormalConfigWithBlogs(),
				Logger:   log.NewNopLogger(),
				instance: new(pastebin.MarkDownImpl),
				prefix:   "test.binacs.space" + "/blog/",
				gitOwner: "binacs",
				gitRepo:  "blog",
				opt:      github.RepositoryContentGetOptions{Ref: "main"},
				client:   github.NewClient(nil),
			},
			args: args{
				uri: "README.md",
			},
			want:    types.Blog{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BlogServiceImpl{
				Config:   tt.fields.Config,
				Logger:   tt.fields.Logger,
				instance: tt.fields.instance,
				prefix:   tt.fields.prefix,
				gitOwner: tt.fields.gitOwner,
				gitRepo:  tt.fields.gitRepo,
				opt:      tt.fields.opt,
				client:   tt.fields.client,
			}
			got, err := bs.URLSearch(tt.args.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("BlogServiceImpl.URLSearch() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				// 	t.Errorf("BlogServiceImpl.URLSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

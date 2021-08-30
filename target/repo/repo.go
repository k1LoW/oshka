package repo

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/k1LoW/oshka/target"
)

var _ target.Target = (*Repo)(nil)

type Repo struct {
	url string
}

func New(repo string) (*Repo, error) {
	u, err := normalize(repo)
	if err != nil {
		return nil, err
	}
	return &Repo{
		url: u,
	}, nil
}

func (r *Repo) Id() string {
	return fmt.Sprintf("repo-%s", target.HashForID([]byte(r.url)))
}

func (r *Repo) Name() string {
	return strings.TrimPrefix(strings.TrimPrefix(r.url, "https://"), "http://")
}

func (r *Repo) Type() string {
	return "repo"
}

func (r *Repo) Dir() string {
	return r.Name()
}

func (r *Repo) Extract(ctx context.Context, dest string) error {
	u := fmt.Sprintf("%s.git", r.url)
	if os.Getenv("GITHUB_TOKEN") != "" {
		_, err := git.PlainClone(dest, false, &git.CloneOptions{
			Auth: &http.BasicAuth{
				Username: "dummy",
				Password: os.Getenv("GITHUB_TOKEN"),
			},
			URL:      u,
			Progress: os.Stdout,
			Depth:    1,
		})
		if err == nil {
			et := new(target.ExtractedTarget)
			if err := et.SetTarget(r, dest); err != nil {
				return err
			}
			if err := et.Put(); err != nil {
				return err
			}

			return nil
		}
		// fallback to the code following here
	}
	_, err := git.PlainClone(dest, false, &git.CloneOptions{
		URL:      u,
		Progress: os.Stdout,
		Depth:    1,
	})
	if err != nil {
		return err
	}

	et := new(target.ExtractedTarget)
	if err := et.SetTarget(r, dest); err != nil {
		return err
	}
	if err := et.Put(); err != nil {
		return err
	}

	return nil
}

func normalize(in string) (string, error) {
	if strings.HasPrefix(in, "git@") {
		return "", fmt.Errorf("SSH URL is not supported:%s", in)
	}
	raw := in
	if !strings.HasPrefix(in, "https://") && !strings.HasPrefix(in, "http://") {
		if strings.Count(in, "/") == 1 && !strings.Contains(in, ".") {
			raw = fmt.Sprintf("https://github.com/%s", in)
		} else {
			raw = fmt.Sprintf("https://%s", in)
		}
	}
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	if strings.Count(u.Path, "/") != 2 {
		return "", fmt.Errorf("invalid repo URL: %s", in)
	}
	return fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, u.Path), nil
}

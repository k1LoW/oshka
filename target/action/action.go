package action

import (
	"context"
	"crypto/md5"
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type Action struct {
	action string
}

func New(action string) (*Action, error) {
	return &Action{
		action: action,
	}, nil
}

func (a *Action) Id() string {
	return fmt.Sprintf("action-%x", md5.Sum([]byte(a.action)))
}

func (a *Action) Name() string {
	return a.action
}

func (a *Action) Type() string {
	return "action"
}

func (a *Action) Dir() string {
	return a.action
}

func (a *Action) Extract(ctx context.Context, dest string) error {
	ownerrepo, tag, err := parse(a.action)
	if err != nil {
		return err
	}
	refName := plumbing.ReferenceName("")
	if tag != "" {
		if strings.HasPrefix(tag, "v") {
			refName = plumbing.ReferenceName(fmt.Sprintf("refs/tags/%s", tag))
		} else {
			refName = plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", tag))
		}
	}

	if os.Getenv("GITHUB_SERVER_URL") != "" && os.Getenv("GITHUB_TOKEN") != "" {
		url := fmt.Sprintf("%s/%s.git", os.Getenv("GITHUB_SERVER_URL"), ownerrepo)
		r, err := git.PlainClone(dest, false, &git.CloneOptions{
			Auth: &http.BasicAuth{
				Username: "dummy",
				Password: os.Getenv("GITHUB_TOKEN"),
			},
			URL:      url,
			Progress: os.Stdout,
		})

		if err == nil {
			w, err := r.Worktree()
			if err != nil {
				return err
			}
			if err := w.Checkout(&git.CheckoutOptions{
				Branch: refName,
			}); err != nil {
				if err := w.Checkout(&git.CheckoutOptions{
					Hash: plumbing.NewHash(tag),
				}); err != nil {
					return err
				}
			}
			return nil
		}
		// fallback to the code following here
	}

	url := fmt.Sprintf("%s/%s.git", "https://github.com", ownerrepo)
	r, err := git.PlainClone(dest, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}
	if err := w.Checkout(&git.CheckoutOptions{
		Branch: refName,
	}); err != nil {
		if err := w.Checkout(&git.CheckoutOptions{
			Hash: plumbing.NewHash(tag),
		}); err != nil {
			return err
		}
	}
	return nil
}

func parse(action string) (owenerrepo string, tag string, err error) {
	if strings.Count(action, "/") == 0 {
		return "", "", fmt.Errorf("invalid action: %s", action)
	}
	if !strings.Contains(action, "@") {
		return "", "", fmt.Errorf("invalid action: %s", action)
	}
	splitted := strings.Split(action, "@")
	splitted2 := strings.Split(splitted[0], "/")
	return strings.Join(splitted2[0:2], "/"), splitted[1], nil
}

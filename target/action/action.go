package action

import (
	"context"
	"crypto/md5"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
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
	ownerrepo, _, tag, branchOrHash, err := parse(a.action)
	if err != nil {
		return err
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
			if tag != "" {
				tagref, err := getTagRef(r, tag)
				if err != nil {
					return err
				}
				if err := w.Checkout(&git.CheckoutOptions{
					Branch: tagref.Name(),
				}); err != nil {
					return err
				}
			} else if branchOrHash != "" {
				if err := w.Checkout(&git.CheckoutOptions{
					Branch: plumbing.NewBranchReferenceName(branchOrHash),
				}); err != nil {
					if err := w.Checkout(&git.CheckoutOptions{
						Hash: plumbing.NewHash(branchOrHash),
					}); err != nil {
						return err
					}
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

	if tag != "" {
		tagref, err := getTagRef(r, tag)
		if err != nil {
			return err
		}
		if err := w.Checkout(&git.CheckoutOptions{
			Branch: tagref.Name(),
		}); err != nil {
			return err
		}
	} else if branchOrHash != "" {
		if err := w.Checkout(&git.CheckoutOptions{
			Branch: plumbing.NewBranchReferenceName(branchOrHash),
		}); err != nil {
			if err := w.Checkout(&git.CheckoutOptions{
				Hash: plumbing.NewHash(branchOrHash),
			}); err != nil {
				return err
			}
		}
	}

	return nil
}

var verRe = regexp.MustCompile(`^v[0-9][0-9.]*$`)

func getTagRef(r *git.Repository, tag string) (*plumbing.Reference, error) {
	tagrefs, err := r.Tags()
	if err != nil {
		return nil, err
	}
	var ref *plumbing.Reference
	if err := tagrefs.ForEach(func(t *plumbing.Reference) error {
		if t.Name().Short() == tag {
			ref = t
			return storer.ErrStop
		}
		if strings.Contains(t.Name().Short(), tag) {
			ref = t
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return ref, nil
}

func parse(action string) (string, string, string, string, error) {
	if strings.Count(action, "/") == 0 {
		return "", "", "", "", fmt.Errorf("invalid action: %s", action)
	}
	if !strings.Contains(action, "@") {
		return "", "", "", "", fmt.Errorf("invalid action: %s", action)
	}
	splitted := strings.Split(action, "@")
	splitted2 := strings.Split(splitted[0], "/")
	ownerrepo := strings.Join(splitted2[0:2], "/")
	path := ""
	if len(splitted2) > 2 {
		path = strings.Join(splitted2[2:], "/")
	}
	tag := ""
	branchOrHash := ""
	if verRe.MatchString(splitted[1]) {
		tag = splitted[1]
	} else {
		branchOrHash = splitted[1]
	}
	return ownerrepo, path, tag, branchOrHash, nil
}

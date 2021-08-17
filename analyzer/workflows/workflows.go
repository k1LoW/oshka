package workflows

import (
	"context"
	"io"
	"io/fs"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/k1LoW/oshka/target"
	"github.com/k1LoW/oshka/target/action"
	"github.com/k1LoW/oshka/target/image"
)

// Workflows is GitHub Actions Workflows file analyzer
type Workflows struct{}

func New() *Workflows {
	return &Workflows{}
}

type WorkflowFile struct {
	Jobs map[string]Job `yaml:"jobs,omitempty"`
}

type Job struct {
	Container interface{}        `yaml:"container,omitempty"`
	Steps     []Step             `yaml:"steps,omitempty"`
	Services  map[string]Service `yaml:"services,omitempty"`
}

type Step struct {
	Uses string `yaml:"uses,omitempty"`
}

type Service struct {
	Image string `yaml:"image"`
}

func (w *Workflows) Analyze(ctx context.Context, fsys fs.FS) ([]target.Target, error) {
	targets := []target.Target{}
	paths, err := fs.Glob(fsys, ".github/workflows/*.yml")
	if err != nil {
		return nil, err
	}
	paths2, err := fs.Glob(fsys, ".github/workflows/*.yaml")
	if err != nil {
		return nil, err
	}
	paths = append(paths, paths2...)
	for _, p := range paths {
		f, err := fsys.Open(p)
		if err != nil {
			return nil, err
		}
		b, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}
		var wf WorkflowFile
		if err := yaml.UnmarshalContext(ctx, b, &wf); err != nil {
			return nil, err
		}
		if len(wf.Jobs) == 0 {
			continue
		}

		for _, j := range wf.Jobs {
			if len(j.Steps) == 0 {
				continue
			}
			var i string
			switch c := j.Container.(type) {
			case string:
				i = c
			case map[string]interface{}:
				ii, ok := c["image"]
				if ok {
					i = ii.(string)
				}
			}
			if i != "" {
				t, err := image.New(i)
				if err != nil {
					return nil, err
				}
				targets = append(targets, t)
			}

			for _, s := range j.Steps {
				if s.Uses == "" {
					continue
				}
				if strings.Contains(s.Uses, "@") {
					t, err := action.New(s.Uses)
					if err != nil {
						return nil, err
					}
					targets = append(targets, t)
				}
			}

			for _, s := range j.Services {
				if s.Image == "" {
					continue
				}
				t, err := image.New(s.Image)
				if err != nil {
					return nil, err
				}
				targets = append(targets, t)
			}
		}
	}
	return targets, nil
}

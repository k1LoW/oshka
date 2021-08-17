package action

import (
	"context"
	"encoding/json"
	"io"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/k1LoW/oshka/target"
	"github.com/k1LoW/oshka/target/dockerfile"
	"github.com/k1LoW/oshka/target/image"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

type Action struct{}

func New() *Action {
	return &Action{}
}

type ActionFile struct {
	Runs Runs `yaml:"runs,omitempty"`
}

type Runs struct {
	Using string `yaml:"using,omitempty"`
	Image string `yaml:"image,omitempty"`
}

func (a *Action) Analyze(ctx context.Context, fsys fs.FS) ([]target.Target, error) {
	targets := []target.Target{}
	f, path, err := openActionYAML(fsys)
	if err != nil {
		return targets, nil
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var af ActionFile
	if err := yaml.UnmarshalContext(ctx, b, &af); err != nil {
		return nil, err
	}
	if af.Runs.Using != "docker" {
		return targets, nil
	}
	if strings.HasPrefix(af.Runs.Image, "docker://") {
		t, err := image.New(strings.TrimPrefix(af.Runs.Image, "docker://"))
		if err != nil {
			return nil, err
		}
		targets = append(targets, t)
	} else {
		df := filepath.Join(path, af.Runs.Image)
		files := map[string][]byte{}

		f, err := fsys.Open(df)
		if err != nil {
			return nil, err
		}
		res, err := parser.Parse(f)
		if err != nil {
			return nil, err
		}

		d := filepath.Dir(df)
		for _, child := range res.AST.Children {
			if child.Next != nil && len(child.Next.Children) > 0 {
				child = child.Next.Children[0]
			}
			if strings.ToUpper(child.Value) != "COPY" {
				continue
			}
			values := []string{}
			for n := child.Next; n != nil; n = n.Next {
				values = append(values, n.Value)
			}
			path := strings.TrimPrefix(filepath.Join(d, values[0]), "/")
			f, err := fsys.Open(path)
			if err != nil {
				return nil, err
			}
			b, err := io.ReadAll(f)
			if err != nil {
				return nil, err
			}
			files[path] = b
		}

		{
			f, err := fsys.Open(df)
			if err != nil {
				return nil, err
			}
			b, err := io.ReadAll(f)
			if err != nil {
				return nil, err
			}
			files[df] = b
		}

		t, err := dockerfile.New(df, files)
		if err != nil {
			return nil, err
		}
		targets = append(targets, t)
	}

	return targets, nil
}

func openActionYAML(fsys fs.FS) (fs.File, string, error) {
	path := ""
	tf, err := fsys.Open(target.ExtractedTargetFile)
	if err == nil {
		et := new(target.ExtractedTarget)
		err := json.NewDecoder(tf).Decode(et)
		if err == nil {
			path = et.ActionYAMLPath
		}
	}
	f, err := fsys.Open(filepath.Join(path, "action.yml"))
	if err != nil {
		f, err = fsys.Open(filepath.Join(path, "action.yaml"))
		if err != nil {
			return nil, "", err
		}
	}
	return f, path, nil
}

package analyzer

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/k1LoW/osfs"
	"github.com/k1LoW/oshka/analyzer/action"
	"github.com/k1LoW/oshka/analyzer/workflows"
	"github.com/k1LoW/oshka/target"
	"github.com/rs/zerolog/log"
)

var (
	_ Analyzer = (*workflows.Workflows)(nil)
	_ Analyzer = (*action.Action)(nil)
)

type Analyzer interface {
	Analyze(ctx context.Context, fsys fs.FS) ([]target.Target, error)
}

func Analyze(ctx context.Context, dirs []string) ([]target.Target, error) {
	tm := map[string]target.Target{}
	analyzers := []Analyzer{action.New(), workflows.New()}

	for _, d := range dirs {
		dir, err := filepath.Abs(d)
		if err != nil {
			return nil, err
		}
		rootfsys := osfs.New()
		fsys, err := fs.Sub(rootfsys, strings.TrimPrefix(dir, "/"))
		if err != nil {
			return nil, err
		}

		for _, a := range analyzers {
			targets, err := a.Analyze(ctx, fsys)
			if err != nil {
				return nil, err
			}
			for _, t := range targets {
				_, ok := tm[t.Id()]
				if ok {
					continue
				}
				log.Info().Msg(fmt.Sprintf("Detect %s %s from %s", t.Type(), t.Name(), dir))
				tm[t.Id()] = t
			}
		}
	}

	targets := []target.Target{}
	for _, t := range tm {
		targets = append(targets, t)
	}
	return targets, nil
}

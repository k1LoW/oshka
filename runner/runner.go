package runner

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/k1LoW/oshka/analyzer"
	"github.com/k1LoW/oshka/executer"
	"github.com/k1LoW/oshka/target"
	"github.com/rs/zerolog/log"
)

type Runner struct {
	targets []target.Target
	e       *executer.Executer
}

func New(targets []target.Target, e *executer.Executer) (*Runner, error) {
	return &Runner{
		targets: targets,
		e:       e,
	}, nil
}

func (r *Runner) Run(ctx context.Context, depth int) error {
	extractRoot := os.TempDir()
	log.Info().Msg(fmt.Sprintf("Create temporary directory for extracting supply chains: %s", extractRoot))
	defer func() {
		_ = os.RemoveAll(extractRoot)
		log.Info().Msg(fmt.Sprintf("Cleanup temporary directory for extracting supply chains: %s", extractRoot))
	}()
	targets := r.targets
	var err error
	i := 0
	for {
		dirs := []string{}
		for _, t := range targets {
			dest := filepath.Join(extractRoot, t.Dir())
			if _, err := os.Stat(dest); err == nil {
				// already extracted
				continue
			}
			log.Info().Msg(fmt.Sprintf("Extract %s %s to %s", t.Type(), t.Name(), dest))
			if err := t.Extract(ctx, dest); err != nil {
				return err
			}
			if err := r.e.Execute(ctx, t, dest); err != nil {
				return err
			}
			dirs = append(dirs, dest)
		}
		i += 1
		if i > depth {
			break
		}
		targets, err = analyzer.Analyze(ctx, dirs)
		if err != nil {
			return err
		}
	}
	return nil
}

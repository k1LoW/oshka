package runner

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/k1LoW/oshka/analyzer"
	"github.com/k1LoW/oshka/executer"
	"github.com/k1LoW/oshka/target"
	"github.com/rs/zerolog/log"
)

type Runner struct {
	e            *executer.Executer
	targetDepths map[string]Depths
}

type Depths []int

func (d Depths) String() string {
	s := []string{}
	for _, i := range d {
		s = append(s, strconv.Itoa(i))
	}
	return strings.Join(s, ", ")
}

func New(e *executer.Executer) (*Runner, error) {
	return &Runner{
		e:            e,
		targetDepths: map[string]Depths{},
	}, nil
}

func (r *Runner) Run(ctx context.Context, targets []target.Target, depth int) error {
	extractRoot := os.TempDir()
	log.Info().Msg(fmt.Sprintf("Create temporary directory for extracting supply chains: %s", extractRoot))
	defer func() {
		_ = os.RemoveAll(extractRoot)
		log.Info().Msg(fmt.Sprintf("Cleanup temporary directory for extracting supply chains: %s", extractRoot))
	}()
	var err error
	i := 0
	for {
		dirs := []string{}
		for _, t := range targets {
			if _, ok := r.targetDepths[t.Id()]; !ok {
				r.targetDepths[t.Id()] = []int{}
			}
			r.targetDepths[t.Id()] = append(r.targetDepths[t.Id()], i)

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

func (r *Runner) TargetDepths(id string) Depths {
	return r.targetDepths[id]
}

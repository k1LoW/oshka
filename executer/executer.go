package executer

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/k1LoW/oshka/target"
	"github.com/rs/zerolog/log"
)

type Executer struct {
	commands [][]string
	results  []*Result
}

type Result struct {
	Command  string
	Dir      string
	Target   target.Target
	Stdout   string
	Stderr   string
	ExitCode int
}

func New(in []string) (*Executer, error) {
	commands := [][]string{}
	for _, s := range in {
		c := strings.Split(s, " ")
		if _, err := exec.LookPath(c[0]); err != nil {
			return nil, err
		}
		commands = append(commands, c)
	}
	return &Executer{
		commands: commands,
		results:  []*Result{},
	}, nil
}

func (e *Executer) Execute(ctx context.Context, t target.Target, dir string) error {
	for _, c := range e.commands {
		r := &Result{
			Command: strings.Join(c, " "),
			Dir:     dir,
			Target:  t,
		}
		log.Info().Msg(fmt.Sprintf("Run `%s` on %s", r.Command, r.Dir))
		cmd := exec.CommandContext(ctx, c[0]) // #nosec G204
		if len(c) > 1 {
			cmd = exec.CommandContext(ctx, c[0], c[1:]...) // #nosec G204
		}
		bufout := new(bytes.Buffer)
		cmd.Stdout = io.MultiWriter(bufout, os.Stdout)
		buferr := new(bytes.Buffer)
		cmd.Stderr = io.MultiWriter(buferr, os.Stderr)
		cmd.Dir = dir
		_ = cmd.Run()
		r.Stdout = bufout.String()
		r.Stderr = buferr.String()
		r.ExitCode = cmd.ProcessState.ExitCode()
		e.results = append(e.results, r)
	}
	return nil
}

func (e *Executer) Results() []*Result {
	return e.results
}

func (e *Executer) ExitCode() int {
	for _, r := range e.results {
		if r.ExitCode != 0 {
			return 1
		}
	}
	return 0
}

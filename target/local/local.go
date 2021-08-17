package local

import (
	"context"
	"fmt"
	"os"

	"github.com/k1LoW/oshka/target"
)

var _ target.Target = (*Local)(nil)

type Local struct {
	dir string
}

func New(dir string) (*Local, error) {
	if _, err := os.Stat(dir); err != nil {
		return nil, err
	}

	return &Local{
		dir: dir,
	}, nil
}

func (l *Local) Id() string {
	return fmt.Sprintf("local-%s", target.HashForID([]byte(l.dir)))
}

func (l *Local) Dir() string {
	return l.Id()
}

func (l *Local) Name() string {
	return l.dir
}

func (l *Local) Type() string {
	return "local"
}

func (l *Local) Extract(ctx context.Context, dest string) error {
	return nil
}

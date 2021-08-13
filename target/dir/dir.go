package dir

import (
	"context"
	"os"
)

type Dir struct {
	dir string
}

func New(dir string) (*Dir, error) {
	if _, err := os.Stat(dir); err != nil {
		return nil, err
	}

	return &Dir{
		dir: dir,
	}, nil
}

func (d *Dir) Id() string {
	return d.dir
}

func (d *Dir) Dir() string {
	return d.dir
}

func (d *Dir) Name() string {
	return d.dir
}

func (d *Dir) Type() string {
	return "dir"
}

func (d *Dir) Extract(ctx context.Context, dest string) error {
	return nil
}

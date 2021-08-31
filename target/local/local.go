package local

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/k1LoW/oshka/target"
	"github.com/otiai10/copy"
)

var _ target.Target = (*Local)(nil)

type Local struct {
	dir  string
	hash string
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

func (l *Local) Hash() string {
	return l.hash
}

func (l *Local) HashType() string {
	return "dir hash"
}

func (l *Local) Extract(ctx context.Context, dest string) error {
	if l.dir == dest {
		return nil
	}

	if err := copy.Copy(l.dir, dest, copy.Options{}); err != nil {
		return err
	}

	hash, err := hashDir(dest)
	if err != nil {
		return err
	}
	l.hash = hash

	et := new(target.ExtractedTarget)
	if err := et.SetTarget(l, dest); err != nil {
		return err
	}
	if err := et.Put(); err != nil {
		return err
	}

	return nil
}

func hashDir(dir string) (string, error) {

	hash := sha256.New()

	if err := filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		io.WriteString(hash, fmt.Sprintf("%s:%d", path, fi.Size()))
		return nil
	}); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

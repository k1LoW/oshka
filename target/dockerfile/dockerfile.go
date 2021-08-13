package dockerfile

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/daemon"
	"github.com/k1LoW/oshka/internal"
)

type Dockerfile struct {
	dockerfile string
	files      map[string][]byte
}

func New(dockerfile string, files map[string][]byte) (*Dockerfile, error) {
	if _, err := exec.LookPath("docker"); err != nil {
		return nil, err
	}
	return &Dockerfile{
		dockerfile: dockerfile,
		files:      files,
	}, nil
}

func (d *Dockerfile) Id() string {
	return fmt.Sprintf("dockerfile-%x", md5.Sum(d.files[d.dockerfile]))
}

func (d *Dockerfile) Name() string {
	return d.Id()
}

func (d *Dockerfile) Type() string {
	return "dockerfile"
}

func (d *Dockerfile) Dir() string {
	return d.Id()
}

func (d *Dockerfile) Extract(ctx context.Context, dest string) error {
	tag := fmt.Sprintf("oshka-tmp-%s", d.Id())
	wd := filepath.Join(os.TempDir(), tag)
	if err := os.MkdirAll(wd, os.ModePerm); err != nil {
		return err
	}
	defer os.RemoveAll(wd)

	for path, b := range d.files {
		p := filepath.Join(wd, path)
		if err := os.MkdirAll(filepath.Dir(p), os.ModePerm); err != nil {
			return err
		}
		if err := os.WriteFile(p, b, os.ModePerm); err != nil {
			return err
		}
	}

	dir := strings.TrimPrefix(filepath.Dir(d.dockerfile), "/")
	if dir == "" {
		dir = "."
	}

	cmd := exec.CommandContext(ctx, "docker", "build", "-t", tag, "-f", d.dockerfile, dir)
	cmd.Dir = wd
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	defer func() {
		_ = exec.Command("docker", "rmi", tag).Run()
	}()

	ref, err := name.ParseReference(tag)
	if err != nil {
		return err
	}
	img, err := daemon.Image(ref)
	if err != nil {
		return err
	}
	r, w := io.Pipe()
	errChan := make(chan error)
	go func() {
		err := internal.Untar(r, dest)
		errChan <- err
	}()
	if err := crane.Export(img, w); err != nil {
		return err
	}
	err = <-errChan
	return err
}

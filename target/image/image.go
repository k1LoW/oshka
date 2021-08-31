package image

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/k1LoW/oshka/internal"
	"github.com/k1LoW/oshka/target"
)

var _ target.Target = (*Image)(nil)

type Image struct {
	image string
	hash  string
}

func New(image string) (*Image, error) {
	return &Image{
		image: image,
	}, nil
}

func (i *Image) Id() string {
	return fmt.Sprintf("image-%s", target.HashForID([]byte(i.image)))
}

func (i *Image) Name() string {
	return i.image
}

func (i *Image) Type() string {
	return "image"
}

func (i *Image) Hash() string {
	return i.hash
}

func (i *Image) HashType() string {
	return "digest"
}

func (i *Image) Dir() string {
	return strings.Replace(i.image, ":", "/", -1)
}

func (i *Image) Extract(ctx context.Context, dest string) error {
	img, err := crane.Pull(i.image, crane.WithContext(ctx))
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

	if err != nil {
		return err
	}

	digest, err := img.Digest()
	if err != nil {
		return err
	}
	i.hash = digest.String()

	et := new(target.ExtractedTarget)
	if err := et.SetTarget(i, dest); err != nil {
		return err
	}
	if err := et.Put(); err != nil {
		return err
	}

	return nil
}

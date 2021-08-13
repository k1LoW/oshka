package target

import (
	"context"

	"github.com/k1LoW/oshka/target/action"
	"github.com/k1LoW/oshka/target/dir"
	"github.com/k1LoW/oshka/target/dockerfile"
	"github.com/k1LoW/oshka/target/image"
)

var (
	_ Target = (*action.Action)(nil)
	_ Target = (*dockerfile.Dockerfile)(nil)
	_ Target = (*image.Image)(nil)
	_ Target = (*dir.Dir)(nil)
)

type Target interface {
	Id() string
	Name() string
	Type() string
	Dir() string
	Extract(ctx context.Context, dest string) error
}

package image

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/k1LoW/oshka/target"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		image   string
		wantErr bool
	}{
		{"alpine:latest", false},
		{"alpine:invalidtag", true},
		{"gcr.io/distroless/static", false},
		{"ghcr.io/k1low/tbls:latest", false},
	}
	root := t.TempDir()
	for _, tt := range tests {
		ctx := context.Background()
		dest := filepath.Join(root, tt.image)
		if err := os.MkdirAll(dest, os.ModePerm); err != nil {
			t.Fatal(err)
		}
		i, err := New(tt.image)
		if err != nil {
			t.Fatal(err)
		}
		gotErr := i.Extract(ctx, dest)
		if gotErr != nil {
			if tt.wantErr {
				continue
			}
			t.Error(gotErr)
		}
		etcDir := filepath.Join(dest, "etc")
		fi, err := os.Stat(etcDir)
		if err != nil {
			t.Error(err)
			continue
		}
		if !fi.IsDir() {
			t.Errorf("%s should be directory", etcDir)
		}

		infoJSON := filepath.Join(dest, target.ExtractedTargetFile)
		if _, err := os.Stat(infoJSON); err != nil {
			t.Error(err)
		}
	}
}

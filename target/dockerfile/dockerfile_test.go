package dockerfile

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/k1LoW/oshka/target"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		dockerfile string
		files      map[string][]byte
		wantErr    bool
	}{
		{
			"Dockerfile",
			map[string][]byte{
				"Dockerfile": []byte(`FROM alpine:latest
RUN touch /etc/test
`),
			},
			false,
		},
	}
	root := t.TempDir()
	for _, tt := range tests {
		ctx := context.Background()
		dest := filepath.Join(root, fmt.Sprintf("%x", sha256.Sum256(tt.files[tt.dockerfile])))
		if err := os.MkdirAll(dest, os.ModePerm); err != nil {
			t.Fatal(err)
		}
		d, err := New("from", tt.dockerfile, tt.files)
		if err != nil {
			t.Fatal(err)
		}
		gotErr := d.Extract(ctx, dest)
		if gotErr != nil {
			if tt.wantErr {
				continue
			}
			t.Error(gotErr)
		}
		testFile := filepath.Join(dest, "etc", "test")
		fi, err := os.Stat(testFile)
		if err != nil {
			t.Error(err)
			continue
		}
		if fi.IsDir() {
			t.Errorf("%s should be file", testFile)
		}

		infoJSON := filepath.Join(dest, target.ExtractedTargetFile)
		if _, err := os.Stat(infoJSON); err != nil {
			t.Error(err)
		}
	}
}

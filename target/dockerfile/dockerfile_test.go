package dockerfile

import (
	"context"
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"testing"
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
		dest := filepath.Join(root, fmt.Sprintf("%x", md5.Sum(tt.files[tt.dockerfile])))
		if err := os.MkdirAll(dest, os.ModePerm); err != nil {
			t.Fatal(err)
		}
		d, err := New(tt.dockerfile, tt.files)
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
	}
}

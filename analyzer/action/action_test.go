package action

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/k1LoW/osfs"
)

func TestAnalize(t *testing.T) {
	tests := []struct {
		path string
		want int
	}{
		//		{"testdata/nodejs", 0},
		//		{"testdata/dockerfile", 1},
		//		{"testdata/dockerimage", 1},
		{"testdata/extracted_target", 1},
	}
	dir := t.TempDir()
	for _, tt := range tests {
		ctx := context.Background()
		fsys := osfs.New()
		d, err := filepath.Abs(tt.path)
		if err != nil {
			t.Fatal(err)
		}
		sub, err := fsys.Sub(strings.TrimPrefix(d, "/"))
		if err != nil {
			t.Fatal(err)
		}
		a := New()
		targets, err := a.Analyze(ctx, sub)
		if err != nil {
			t.Error(err)
		}
		got := len(targets)
		if got != tt.want {
			t.Errorf("got %v\nwant %v", got, tt.want)
		}
		for _, tgt := range targets {
			dest := filepath.Join(dir, tgt.Dir())
			if err := os.MkdirAll(dest, os.ModePerm); err != nil {
				t.Fatal(err)
			}
			if err := tgt.Extract(ctx, dest); err != nil {
				t.Error(err)
			}
		}
	}
}

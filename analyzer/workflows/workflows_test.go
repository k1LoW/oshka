package workflows

import (
	"context"
	"path/filepath"
	"strings"
	"testing"

	"github.com/k1LoW/osfs"
)

func TestAnalyze(t *testing.T) {
	ctx := context.Background()
	fsys := osfs.New()
	d, err := filepath.Abs("testdata")
	if err != nil {
		t.Fatal(err)
	}
	sub, err := fsys.Sub(strings.TrimPrefix(d, "/"))
	if err != nil {
		t.Fatal(err)
	}

	w := New()
	targets, err := w.Analyze(ctx, sub)
	if err != nil {
		t.Fatal(err)
	}

	got := len(targets)
	if want := 8; got != want {
		t.Errorf("got %v\nwant %v", got, want)
	}
}

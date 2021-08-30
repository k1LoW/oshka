package repo

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/k1LoW/oshka/target"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		repo    string
		wantErr bool
	}{
		{"actions/cache", false},
		{"github.com/does-not/exists", true},
	}
	root := t.TempDir()
	for _, tt := range tests {
		ctx := context.Background()
		r, err := New(tt.repo)
		if err != nil {
			t.Fatal(err)
		}
		dest := filepath.Join(root, r.Dir())
		if err := os.MkdirAll(dest, os.ModePerm); err != nil {
			t.Fatal(err)
		}
		gotErr := r.Extract(ctx, dest)
		if gotErr != nil {
			if tt.wantErr {
				continue
			}
			t.Error(gotErr)
		}
		gitDir := filepath.Join(dest, ".git")
		fi, err := os.Stat(gitDir)
		if err != nil {
			t.Error(err)
			continue
		}
		if !fi.IsDir() {
			t.Errorf("%s should be directory", gitDir)
		}

		infoJSON := filepath.Join(dest, target.ExtractedTargetFile)
		if _, err := os.Stat(infoJSON); err != nil {
			t.Error(err)
		}
	}
}

func TestName(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"https://github.com/k1LoW/oshka", "github.com/k1LoW/oshka"},
		{"github.com/k1LoW/oshka", "github.com/k1LoW/oshka"},
		{"git.company.com/k1LoW/oshka", "git.company.com/k1LoW/oshka"},
		{"k1LoW/oshka", "github.com/k1LoW/oshka"},
	}
	for _, tt := range tests {
		r, err := New(tt.in)
		if err != nil {
			t.Fatal(err)
		}
		got := r.Name()
		if got != tt.want {
			t.Errorf("got %v\nwant %v", got, tt.want)
		}
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		in      string
		want    string
		wantErr bool
	}{
		{"https://github.com/k1LoW/oshka", "https://github.com/k1LoW/oshka", false},
		{"github.com/k1LoW/oshka", "https://github.com/k1LoW/oshka", false},
		{"git.company.com/k1LoW/oshka", "https://git.company.com/k1LoW/oshka", false},
		{"k1LoW/oshka", "https://github.com/k1LoW/oshka", false},
		{"git@github.com:k1LoW/oshka.git", "", true},
	}
	for _, tt := range tests {
		got, err := normalize(tt.in)
		if err != nil {
			if !tt.wantErr {
				t.Error(err)
			}
			continue
		}
		if got != tt.want {
			t.Errorf("got %v\nwant %v", got, tt.want)
		}
	}
}

package action

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		action  string
		wantErr bool
	}{
		{"actions/cache@v2", false},
		{"actions/cache@main", false},
		{"github/codeql-action/init@v1", false},
		{"ruby/setup-ruby@473e4d8fe5dd94ee328fdfca9f8c9c7afc9dae5e", false},
		{"does-not/exists@v5", true},
	}
	root := t.TempDir()
	for _, tt := range tests {
		ctx := context.Background()
		dest := filepath.Join(root, tt.action)
		if err := os.MkdirAll(dest, os.ModePerm); err != nil {
			t.Fatal(err)
		}
		a, err := New(tt.action)
		if err != nil {
			t.Fatal(err)
		}
		gotErr := a.Extract(ctx, dest)
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
	}
}

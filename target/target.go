package target

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/goccy/go-json"
)

type Target interface {
	Id() string
	Name() string
	Type() string
	Dir() string
	Extract(ctx context.Context, dest string) error
}

const ExtractedTargetFile = ".extracted_target.json"

type ExtractedTarget struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	ExtractedDir string    `json:"extracted_dir"`
	ExtractedAt  time.Time `json:"extracted_at"`

	ActionYAMLPath string `json:"action_yaml_path"`
}

func (e *ExtractedTarget) SetTarget(t Target, dest string) error {
	e.Id = t.Id()
	e.Name = t.Name()
	e.Type = t.Type()
	e.ExtractedDir = dest
	e.ExtractedAt = time.Now()
	return nil
}

func (e *ExtractedTarget) Put() error {
	f, err := os.Create(filepath.Join(e.ExtractedDir, ExtractedTargetFile))
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	return json.NewEncoder(f).Encode(e)
}

func HashForID(b []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(b))[:7]
}

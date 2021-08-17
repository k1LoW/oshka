package internal

import (
	"archive/tar"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var maxCopySize int64 = 1024 * 1024 * 1024 * 5 // 5GB

func Untar(r io.Reader, dest string) error {
	tr := tar.NewReader(r)
	for {
		f, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if !validRelPath(f.Name) {
			return fmt.Errorf("invalid name: %s", f.Name)
		}
		rel := filepath.FromSlash(f.Name)
		path := filepath.Join(dest, rel)

		fi := f.FileInfo()
		mode := fi.Mode()
		if mode.IsDir() {
			if err := os.MkdirAll(path, 0750); err != nil {
				return err
			}
		} else {
			if err := os.MkdirAll(filepath.Dir(path), 0750); err != nil {
				return err
			}
			of, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, mode.Perm())
			if err != nil {
				return err
			}
			w, err := io.CopyN(of, tr, maxCopySize)
			if !errors.Is(err, io.EOF) {
				return err
			}
			if w > maxCopySize {
				return fmt.Errorf("copy size error: %d > %d", w, maxCopySize)
			}
			if err := of.Close(); err != nil {
				return err
			}
		}
	}
	return nil
}

func validRelPath(p string) bool {
	if p == "" || strings.Contains(p, `\`) || strings.HasPrefix(p, "/") || strings.Contains(p, "../") {
		return false
	}
	return true
}

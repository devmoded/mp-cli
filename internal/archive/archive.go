package archive

import (
	"fmt"
	"path/filepath"
	"strings"
)

type Archive interface {
	Inspect() (Inspect, error)
	ExtractAll(dest string) error
	ExtractEntry(path, dest string) error
}

func Open(path string) (Archive, error) {
	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".zip":
		return NewZipArchive(path), nil
	default:
		return nil, fmt.Errorf("unknown archive extension")
	}
}

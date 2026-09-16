package archive

import (
	"encoding/json"
	"path/filepath"
	"strings"
)

type Metadata struct {
	Name string `json:"name,omitempty"`
}

func (m Metadata) Sprint() string {
	b, _ := json.Marshal(m)
	return string(b)
}

type Entry struct {
	Path string `json:"path,omitempty"`
}

type Archive interface {
	Inspect() (Metadata, []Entry, error)
	ExtractAll(dest string) error
	ExtractEntry(path, dest string) error
}

func Open(path string) Archive {
	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".zip":
		return NewZipArchive(path)
	}
	return nil
}

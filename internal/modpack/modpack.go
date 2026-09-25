package modpack

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/devmoded/mp-cli/internal/data"
	"github.com/devmoded/mp-cli/internal/modpack/basic"
)

type Modpack interface {
	InspectContent() ([]data.Entry, error)
	InspectMetadata() (data.Metadata, error)
	Extract(paths []string, dest string) error
}

func Open(path string) (Modpack, error) {
	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".zip":
		return basic.NewBasicModpack(path), nil
	default:
		return nil, fmt.Errorf("unknown archive extension")
	}
}

func Inspect(m Modpack) (data.Inspect, error) {
	entries, err := m.InspectContent()
	if err != nil {
		return data.Inspect{}, err
	}

	meta, err := m.InspectMetadata()
	if err != nil {
		return data.Inspect{}, err
	}
	return data.Inspect{
		Entries:  entries,
		Metadata: meta,
	}, nil
}

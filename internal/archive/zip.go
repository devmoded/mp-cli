package archive

import (
	"archive/zip"
	"encoding/json"
)

type ZipArchive struct {
	path string
}

func NewZipArchive(path string) ZipArchive {
	return ZipArchive{
		path: path,
	}
}

func readMetadata(f *zip.File) (Metadata, error) {
	rc, err := f.Open()
	if err != nil {
		return Metadata{}, err
	}
	defer rc.Close()

	var meta Metadata
	if err := json.NewDecoder(rc).Decode(&meta); err != nil {
		return Metadata{}, err
	}

	return meta, nil
}

func (za ZipArchive) Inspect() (Inspect, error) {
	r, err := zip.OpenReader(za.path)
	if err != nil {
		return Inspect{}, err
	}
	defer r.Close()

	var entries []Entry
	var meta Metadata

	for _, f := range r.File {
		entries = append(entries, Entry{Path: f.Name})

		if f.Name == "info.json" {
			meta, err = readMetadata(f)
			if err != nil {
				return Inspect{}, err
			}
		}
	}
	return Inspect{Metadata: meta, Entries: entries}, nil
}

// TODO: Сделать распаковку всего архива
func (za ZipArchive) ExtractAll(dest string) error {
	return nil
}

// TODO: Сделать распаковку указанного файла или каталога
func (za ZipArchive) ExtractEntry(path, dest string) error {
	return nil
}

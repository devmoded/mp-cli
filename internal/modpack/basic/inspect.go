package basic

import (
	"archive/zip"
	"encoding/json"

	"github.com/devmoded/mp-cli/internal/data"
)

const MetadataFilename = "info.json"

func readMetadata(f *zip.File) (data.Metadata, error) {
	rc, err := f.Open()
	if err != nil {
		return data.Metadata{}, err
	}
	defer rc.Close()

	var meta data.Metadata
	if err := json.NewDecoder(rc).Decode(&meta); err != nil {
		return data.Metadata{}, err
	}

	return meta, nil
}

func (bm BasicModpack) InspectMetadata() (data.Metadata, error) {
	r, err := zip.OpenReader(bm.path)
	if err != nil {
		return data.Metadata{}, err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name == MetadataFilename {
			return readMetadata(f)
		}
	}
	return data.Metadata{}, nil
}

func (bm BasicModpack) InspectContent() ([]data.Entry, error) {
	r, err := zip.OpenReader(bm.path)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var entries []data.Entry

	for _, f := range r.File {
		entries = append(entries, data.Entry{Path: f.Name})
	}
	return entries, nil
}

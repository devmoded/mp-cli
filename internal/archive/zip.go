package archive

import (
	"archive/zip"
)

type ZipArchive struct {
	path string
}

func NewZipArchive(path string) ZipArchive {
	return ZipArchive{
		path: path,
	}
}

func (za ZipArchive) Inspect() (Metadata, []Entry, error) {
	r, err := zip.OpenReader(za.path)
	if err != nil {
		return Metadata{}, nil, err
	}
	defer r.Close()

	var entries []Entry
	// TODO: Добавить извлечение информации из info.json,
	// и потом аналогичное для mrpack
	var meta Metadata

	for _, f := range r.File {
		entries = append(entries, Entry{Path: f.Name})
	}
	return meta, entries, nil
}

// TODO: Сделать распаковку всего архива
func (za ZipArchive) ExtractAll(dest string) error {
	return nil
}

// TODO: Сделать распаковку указанного файла или каталога
func (za ZipArchive) ExtractEntry(path, dest string) error {
	return nil
}

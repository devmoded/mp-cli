package archive

import (
	"strings"
)

func addField(b *strings.Builder, prefix, separator, value string) {
	if value != "" {
		b.WriteString(prefix)
		b.WriteString(separator)
		b.WriteString(value)
		b.WriteByte('\n')
	}
}

type Inspect struct {
	Metadata Metadata `json:"metadata"`
	Entries  []Entry  `json:"entries"`
}

func (i Inspect) String() string {
	var b strings.Builder
	var e strings.Builder

	for _, entry := range i.Entries {
		addField(&e, "", "  ", entry.Path)
	}

	addField(&b, "Metadata:", "\n", i.Metadata.String())
	addField(&b, "Content:", "\n", strings.TrimSuffix(e.String(), "\n"))

	return strings.TrimSuffix(b.String(), "\n")
}

type Metadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (m Metadata) String() string {
	var b strings.Builder

	addField(&b, "  Name", ": ", m.Name)
	addField(&b, "  Description", ": ", m.Description)
	return strings.TrimSuffix(b.String(), "\n")
}

type Entry struct {
	Path string `json:"path,omitempty"`
}

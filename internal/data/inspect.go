package data

import (
	"strings"
)

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

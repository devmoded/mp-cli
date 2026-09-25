package data

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

type Entry struct {
	Path string `json:"path,omitempty"`
}

package data

import (
	"strings"
)

type Metadata struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	Version          string `json:"version"`
	GameVersion      string `json:"game_version"`
	ModLoader        string `json:"modloader"`
	ModLoaderVersion string `json:"modloader_version"`
}

func (m Metadata) String() string {
	var b strings.Builder

	addField(&b, "  Name", ": ", m.Name)
	addField(&b, "  Description", ": ", m.Description)
	addField(&b, "  Version", ": ", m.Version)
	addField(&b, "  Game version", ": ", m.GameVersion)
	addField(&b, "  Modloader", ": ", m.ModLoader)
	addField(&b, "  Modloader version", ": ", m.ModLoaderVersion)

	return strings.TrimSuffix(b.String(), "\n")
}

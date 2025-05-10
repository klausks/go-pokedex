package model

import (
	"fmt"
	"strings"
)

type Pokemon struct {
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Name           string `json:"name"`
	Order          int    `json:"order"`
	Species        struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`

	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func (p *Pokemon) Info() string {
	text := `
	Name: %s
	Height: %d
	Weight: %d
	%s
	%s
	`
	return fmt.Sprintf(text, p.Name, p.Height, p.Weight, p.statsInfo(), p.typesInfo())
}

func (p *Pokemon) typesInfo() string {
	header := "Types:\n"
	types := make([]string, len(p.Types))
	for i, typ := range p.Types {
		types[i] = fmt.Sprintf("\t-%s", typ.Type.Name)
	}
	return header + strings.Join(types, "\n") + "\n"
}

func (p *Pokemon) statsInfo() string {
	header := "Stats:\n"
	stats := make([]string, len(p.Stats))
	for i, stat := range p.Stats {
		stats[i] = fmt.Sprintf("\t-%s: %d", stat.Stat.Name, stat.BaseStat)
	}
	return header + strings.Join(stats, "\n") + "\n"
}

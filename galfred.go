package galfred

import (
	"encoding/json"
	"fmt"
)

type Icon struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type Text struct {
	Copy      string `json:"copy,omitempty"`
	LargeType string `json:"largetype,omitempty"`
}

type Mod struct {
	valid    string `json:"valid"`
	arg      string `json:"arg"`
	subtitle string `json:"subtitle"`
}

type Item struct {
	Uid          string         `json:"uid,omitempty"`
	Type         string         `json:"type,omitempty"`
	Title        string         `json:"title"`
	Subtitle     string         `json:"subtitle,omitempty"`
	Arg          string         `json:"arg"`
	AutoComplete string         `json:"autocomplete"`
	Valid        bool           `json:"valid,omitempty"`
	Icon         Icon           `json:"icon"`
	Text         Text           `json:"text,omitempty"`
	QuickLookUrl string         `json:"quicklookurl,omitempty"`
	Mods         map[string]Mod `json:"mods,omitempty"`
}

type AlfredTemplate struct {
	Items []Item `json:"items"`
}

func (a *AlfredTemplate) appendItem(item *Item) []Item {
	a.Items = append(a.Items, *item)
	return a.Items
}

func (a *AlfredTemplate) append(title string, subtitle string, arg string, autoComplete string, iconType string, iconPath string) []Item {
	item := Item{
		Uid:          "",
		Type:         "",
		Title:        title,
		Subtitle:     subtitle,
		Arg:          arg,
		AutoComplete: autoComplete,
		Icon: Icon{
			Type: iconType,
			Path: iconPath,
		},
	}
	a.Items = append(a.Items, item)
	return a.Items
}

func NewTemplate() *AlfredTemplate {
	a := &AlfredTemplate{}
	return a
}

func (a *AlfredTemplate) toJson() string {
	json, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(json)
}

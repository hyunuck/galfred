package galfred

import (
	"encoding/json"
	"crypto/rand"
	"fmt"
	"io"
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

func (a *AlfredTemplate) AppendItem(item *Item) []Item {
	a.Items = append(a.Items, *item)
	return a.Items
}

func (a *AlfredTemplate) Append(title string, subtitle string, arg string, autoComplete string, iconType string, iconPath string) []Item {

	uuid, err := newUUID()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	item := Item{
		Uid:          uuid,
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

func (a *AlfredTemplate) AppendTitle(title string, arg string) []Item {

	uuid, err := newUUID()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	item := Item{
		Uid:          uuid,
		Title:        title,
		Arg:          arg,
	}
	a.Items = append(a.Items, item)
	return a.Items
}

func NewTemplate() *AlfredTemplate {
	a := &AlfredTemplate{}
	return a
}

func (a *AlfredTemplate) ToJson() string {
	json, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(json)
}


// newUUID generates a random UUID according to RFC 4122
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

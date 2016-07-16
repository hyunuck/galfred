package galfred

import (
	"testing"
	"fmt"
)

type testdata struct {
	title        string
	subtitle     string
	arg          string
	autoComplete string
	iconType     string
	iconPath     string
}

func TestNewTemplate(t *testing.T) {

	template := NewTemplate()

	td := []testdata{
		testdata{
			"title0",
			"subtitle0",
			"arg0",
			"auto_complete0",
			"icon_type0",
			"icon_path0",
		},
		testdata{
			"title1",
			"subtitle1",
			"arg1",
			"auto_complete1",
			"icon_type1",
			"icon_path1",
		},
	}

	for _, v := range td {
		template.append(v.title, v.subtitle, v.arg, v.autoComplete, v.iconType, v.iconPath)
	}

	t.Log("Test containing correct data.")
	{

		for i, v := range td {
			if v.title == template.Items[i].Title &&
				v.subtitle == template.Items[i].Subtitle &&
				v.arg == template.Items[i].Arg &&
				v.autoComplete == template.Items[i].AutoComplete &&
				v.iconType == template.Items[i].Icon.Type &&
				v.iconPath == template.Items[i].Icon.Path {
				continue
			} else {
				t.Fatal("Template object does not containg correcct data.")
			}
		}
	}

	t.Log("Test jon marshaling")
	{
		json := template.toJson()
		if "" == json {
			t.Fatal("Creating json failed.")
		}
	}

}

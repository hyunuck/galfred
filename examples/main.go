package main

import (
	"fmt"
    "github.com/hyunuck/galfred"
)

type testdata struct {
    title        string
    subtitle     string
    arg          string
    autoComplete string
    iconType     string
    iconPath     string
}

func main() {

    template := galfred.NewTemplate()

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
        template.AppendTitle(v.title, v.title)
    }

    fmt.Println(template.ToJson())

}

## Galfred

Alfred workflow script filter format builder wrtten in Go.   

## Getting started

```sh
$ go get github.com/hyunuck/galfred
```


## Usage 

```go

import "github.com/hyunuck/galfred"

template := galfred.NewTemplate()
template.Append(v.title, v.subtitle, v.arg, v.autoComplete, v.iconType, v.iconPath)
template.toJson()
```
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/wayneashleyberry/css-color?tab=doc)
![go](https://github.com/wayneashleyberry/css-color/workflows/go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/css-color)](https://goreportcard.com/report/github.com/wayneashleyberry/css-color)

> Convert CSS color values into native Go [image/color](https://golang.org/pkg/image/color) values.

The parser currently supports almost all values allowed by [the css spec](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value). Inlcuding 3, 4, 6 and 8 digit hex, rgb, rgba, hsl, hsla and keywords.

Read more on [pkg.go.dev](https://pkg.go.dev/github.com/wayneashleyberry/css-color?tab=doc).

### Installation

```sh
go get github.com/wayneashleyberry/css-color
```

### Usage

```go
package main

import parser "github.com/wayneashleyberry/css-color"

func main() {
    p := parser.New()
    col, _ := p.Convert("#bada55")
    fmt.Println(col)
}
```

[Try the above example in The Go Playground.](https://play.golang.org/p/NYFAJ7B-9D3)

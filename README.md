# Gin

[![Doc](https://pkg.go.dev/badge/github.com/gradii/gin)](https://pkg.go.dev/github.com/gradii/gin)
[![Go](https://img.shields.io/github/go-mod/go-version/gradii/gin)](https://go.dev/)
[![Release](https://img.shields.io/github/release/gradii/gin.svg)](https://github.com/gradii/gin/releases)
[![Test](https://github.com/gradii/gin/actions/workflows/test.yml/badge.svg)](https://github.com/gradii/gin/actions)
[![Report Card](https://goreportcard.com/badge/github.com/gradii/gin)](https://goreportcard.com/report/github.com/gradii/gin)
[![Codecov](https://codecov.io/gh/gradii/gin/branch/main/graph/badge.svg)](https://codecov.io/gh/gradii/gin)
![License](https://img.shields.io/github/license/gradii/gin)

Gin HTTP driver for `github.com/gradii/framework`, a fork of
[goravel/gin](https://github.com/goravel/gin) published as `github.com/gradii/gin`.

## Install

Run the command below in your project to install the package automatically:

```
./artisan package:install github.com/gradii/gin
```

Or check [the setup file](./setup/setup.go) to install the package manually.

## Configuration

You can define the `template` configuration. If omitted, `DefaultTemplate()` is used automatically as a fallback, which loads views from `resources/views` and any registered package views.

You can provide a custom template configuration in two forms:

- **`func() (render.HTMLRender, error)`** — a callback that returns a custom HTML renderer (e.g. to configure custom delimiters or a FuncMap).
- **`render.HTMLRender`** — a pre-built renderer instance.

**Custom example:**

```go
import (
    "html/template"

    "github.com/gin-gonic/gin/render"
    "github.com/gradii/gin"
)

"template": func() (render.HTMLRender, error) {
    return gin.NewTemplate(gin.RenderOptions{
        Delims:  &gin.Delims{Left: "{[", Right: "]}"},
        FuncMap: template.FuncMap{
            "upper": strings.ToUpper,
        },
    })
},
```

## Testing

Run command below to run test:

```
go test ./...
```

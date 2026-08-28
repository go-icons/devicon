# devicon

[![ci](https://github.com/go-icons/devicon/actions/workflows/ci.yml/badge.svg)](https://github.com/go-icons/devicon/actions/workflows/ci.yml)
![coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-icons/devicon.svg)](https://pkg.go.dev/github.com/go-icons/devicon)
[![License](https://img.shields.io/badge/license-BSD--3--Clause-blue.svg)](LICENSE)

Programming-language and tool logos from the
[Devicon](https://github.com/devicons/devicon) icon set (by konpa and
contributors, MIT), as embedded SVG documents keyed by file name — for pure-Go
UIs that render their own icons.

```go
import "github.com/go-icons/devicon"

svg := devicon.Icon("main.go")  // the Go language logo, as an SVG string
dir := devicon.Folder()         // a neutral folder glyph
```

Unlike a full file-type theme, Devicon is a set of **language/tool logos**.
`Icon(filename)` maps a source file's extension to the logo of the language it
is written in (`.go` → Go, `.py` → Python, `.rs` → Rust, `.tex` → LaTeX …),
matching by exact base name first (`go.mod`, `package.json`), then by extension.
Anything with no language logo falls back to a neutral document glyph, and
`Folder()` returns a neutral folder glyph — Devicon ships no generic file or
folder icon, so those two glyphs are minimal originals.

It is a **data package**: it returns SVG strings and draws nothing. A renderer
such as [go-widgets/toolkit](https://github.com/go-widgets/toolkit)'s `SVGIcon`
turns the SVG into a drawn glyph. Every embedded logo is verified to rasterise
under the [go-gfx/gfx](https://github.com/go-gfx/gfx) SVG renderer.

## Licence

The Go code and the two neutral fallback glyphs are BSD-3-Clause (`LICENSE`).
The embedded Devicon logo artwork is MIT, © 2015 konpa (`DEVICON-LICENSE`) —
redistributed unmodified.

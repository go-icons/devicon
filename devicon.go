// Copyright (c) 2026 the go-widgets authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

// Package devicon serves programming-language and tool logos from the Devicon
// icon set (by konpa and contributors, MIT — see DEVICON-LICENSE), as SVG
// documents keyed by file name.
//
// Unlike a full file-type theme, Devicon is a set of colourful language/tool
// logos: [Icon] maps a source file's extension to the logo of the language it
// is written in, falling back to a neutral document glyph for anything with no
// language logo. [Folder] returns a neutral folder glyph. The document and
// folder glyphs are minimal originals (Devicon ships no generic file or folder
// icon).
//
// It is a data package: a curated subset of the logos is embedded, and the API
// returns SVG strings. A renderer such as go-widgets/toolkit's SVGIcon turns
// the returned SVG into a drawn glyph; this package draws nothing itself.
package devicon

import (
	"embed"
	"path"
	"strings"
)

// Name is the human label a picker shows for this pack.
const Name = "Devicon"

//go:embed svg/*.svg
var files embed.FS

// byName maps a lower-cased base file name to a logo, for names that carry more
// meaning than their extension (or have none).
var byName = map[string]string{
	"go.mod":       "go",
	"go.sum":       "go",
	"package.json": "json",
}

// byExt maps a lower-cased source extension (with the dot) to the logo of the
// language it is written in.
var byExt = map[string]string{
	".go": "go",
	".py": "python",
	".rs": "rust",
	".js": "javascript", ".mjs": "javascript", ".cjs": "javascript", ".jsx": "javascript",
	".ts": "typescript", ".tsx": "typescript",
	".c": "c", ".h": "c",
	".cpp": "cplusplus", ".cc": "cplusplus", ".cxx": "cplusplus", ".hpp": "cplusplus", ".hh": "cplusplus",
	".java": "java",
	".rb":   "ruby",
	".php":  "php",
	".cs":   "csharp",
	".sh":   "bash", ".bash": "bash", ".zsh": "bash",
	".lua":  "lua",
	".html": "html5", ".htm": "html5",
	".css":  "css3",
	".json": "json",
	".md":   "markdown", ".markdown": "markdown",
	".tex": "latex", ".sty": "latex", ".cls": "latex",
	".yml": "yaml", ".yaml": "yaml",
	".xml": "xml",
	".pl":  "perl", ".pm": "perl",
	".swift": "swift",
	".kt":    "kotlin", ".kts": "kotlin",
	".scala": "scala", ".sc": "scala",
	".hs": "haskell",
	".ex": "elixir", ".exs": "elixir",
	".dart": "dart",
	".r":    "r",
}

// Icon returns the SVG document for the file named filename (any path — only the
// base name matters), matching by exact name first, then by source extension,
// then a neutral document glyph. It never returns "" for a normal name: the
// default glyph is embedded.
func Icon(filename string) string {
	base := strings.ToLower(path.Base(filename))
	if ic, ok := byName[base]; ok {
		return read(ic)
	}
	if ic, ok := byExt[strings.ToLower(path.Ext(base))]; ok {
		return read(ic)
	}
	return read("default_file")
}

// Folder returns a neutral SVG glyph for a directory.
func Folder() string { return read("default_folder") }

// read returns the embedded SVG for a logo name, or "" when absent.
func read(name string) string {
	b, err := files.ReadFile("svg/" + name + ".svg")
	if err != nil {
		return ""
	}
	return string(b)
}

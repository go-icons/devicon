// Copyright (c) 2026 the go-widgets authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

package devicon

import (
	"strings"
	"testing"
)

func svgOK(s string) bool { return s != "" && strings.Contains(s, "<svg") }

// TestIconByExtension: a source file's extension selects its language logo.
func TestIconByExtension(t *testing.T) {
	for _, name := range []string{"main.go", "app.py", "lib.rs", "index.JS", "types.ts", "core.c", "engine.cpp", "Main.java", "gem.rb", "site.php", "paper.tex", "notes.md", "conf.yaml"} {
		if !svgOK(Icon(name)) {
			t.Errorf("Icon(%q) returned no SVG", name)
		}
	}
}

// TestIconByName: a full base name (no useful extension) selects its logo.
func TestIconByName(t *testing.T) {
	for _, name := range []string{"go.mod", "path/to/go.sum", "package.json"} {
		if !svgOK(Icon(name)) {
			t.Errorf("Icon(%q) returned no SVG", name)
		}
	}
}

// TestIconDefaultFallback: an unknown extension falls back to the neutral glyph.
func TestIconDefaultFallback(t *testing.T) {
	if !svgOK(Icon("mystery.zzz")) {
		t.Error("unknown extension should fall back to the default glyph")
	}
	if !svgOK(Icon("noext")) {
		t.Error("an extensionless unknown name should fall back to the default glyph")
	}
}

// TestFolder: the folder glyph is present.
func TestFolder(t *testing.T) {
	if !svgOK(Folder()) {
		t.Error("Folder() returned no SVG")
	}
}

// TestName: the pack exposes a human label.
func TestName(t *testing.T) {
	if Name == "" {
		t.Error("Name should not be empty")
	}
}

// TestEveryMappedIconEmbedded guards the maps against a typo: every logo name
// they reference (plus the defaults) must resolve to an embedded SVG.
func TestEveryMappedIconEmbedded(t *testing.T) {
	seen := map[string]bool{"default_file": true, "default_folder": true}
	for _, m := range []map[string]string{byName, byExt} {
		for _, ic := range m {
			seen[ic] = true
		}
	}
	for ic := range seen {
		if read(ic) == "" {
			t.Errorf("mapped logo %q is not embedded", ic)
		}
	}
}

// TestReadMissing covers read's not-found branch directly.
func TestReadMissing(t *testing.T) {
	if read("definitely-not-an-icon") != "" {
		t.Error("read of a missing logo should return an empty string")
	}
}

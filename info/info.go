// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package debug exports debug information for gopls.
package info

import (
	"context"
	"fmt"
	"io"
	"reflect"
	"runtime/debug"
	"strings"
)

type PrintMode int

const (
	PlainText = PrintMode(iota)
	Markdown
	HTML
)

// Version is a manually-updated mechanism for tracking versions.
const Version = "master"

// ServerVersion is the format used by gopls to report its version to the
// client. This format is structured so that the client can parse it easily.
type ServerVersion struct {
	Module
	Deps []*Module `json:"deps,omitempty"`
}

type Module struct {
	ModuleVersion
	Replace *ModuleVersion `json:"replace,omitempty"`
}

type ModuleVersion struct {
	Path    string `json:"path,omitempty"`
	Version string `json:"version,omitempty"`
	Sum     string `json:"sum,omitempty"`
}

// VersionInfo returns the build info for the gopls process. If it was not
// built in module mode, we return a GOPATH-specific message with the
// hardcoded version.
func VersionInfo() *ServerVersion {
	if info, ok := debug.ReadBuildInfo(); ok {
		return getVersion(info)
	}
	path := "gopls, built in GOPATH mode"
	return &ServerVersion{
		Module: Module{
			ModuleVersion: ModuleVersion{
				Path:    path,
				Version: Version,
			},
		},
	}
}

func getVersion(info *debug.BuildInfo) *ServerVersion {
	serverVersion := ServerVersion{
		Module: Module{
			ModuleVersion: ModuleVersion{
				Path:    info.Main.Path,
				Version: info.Main.Version,
				Sum:     info.Main.Sum,
			},
		},
	}
	for _, d := range info.Deps {
		m := &Module{
			ModuleVersion: ModuleVersion{
				Path:    d.Path,
				Version: d.Version,
				Sum:     d.Sum,
			},
		}
		if d.Replace != nil {
			m.Replace = &ModuleVersion{
				Path:    d.Replace.Path,
				Version: d.Replace.Version,
			}
		}
		serverVersion.Deps = append(serverVersion.Deps, m)
	}
	return &serverVersion
}

// PrintVersionInfo writes version information to w, using the output format
// specified by mode. verbose controls whether additional information is
// written, including section headers.
func PrintVersionInfo(ctx context.Context, w io.Writer, verbose bool, mode PrintMode) {
	info := VersionInfo()
	if !verbose {
		printBuildInfo(w, info, false, mode)
		return
	}
	section(w, mode, "Build info", func() {
		printBuildInfo(w, info, true, mode)
	})
}

func section(w io.Writer, mode PrintMode, title string, body func()) {
	switch mode {
	case PlainText:
		fmt.Fprintln(w, title)
		fmt.Fprintln(w, strings.Repeat("-", len(title)))
		body()
	case Markdown:
		fmt.Fprintf(w, "#### %s\n\n```\n", title)
		body()
		fmt.Fprintf(w, "```\n")
	case HTML:
		fmt.Fprintf(w, "<h3>%s</h3>\n<pre>\n", title)
		body()
		fmt.Fprint(w, "</pre>\n")
	}
}

func printBuildInfo(w io.Writer, info *ServerVersion, verbose bool, mode PrintMode) {
	fmt.Fprintf(w, "%v %v\n", info.Path, Version)
	printModuleInfo(w, &info.Module, mode)
	if !verbose {
		return
	}
	for _, dep := range info.Deps {
		printModuleInfo(w, dep, mode)
	}
}

func printModuleInfo(w io.Writer, m *Module, mode PrintMode) {
	fmt.Fprintf(w, "    %s@%s", m.Path, m.Version)
	if m.Sum != "" {
		fmt.Fprintf(w, " %s", m.Sum)
	}
	if m.Replace != nil {
		fmt.Fprintf(w, " => %v", m.Replace.Path)
	}
	fmt.Fprintf(w, "\n")
}

type field struct {
	index []int
}

var fields []field

// find all the options. The presumption is that the Options are nested structs
// and that pointers don't need to be dereferenced
func swalk(t reflect.Type, ix []int, indent string) {
	switch t.Kind() {
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			fld := t.Field(i)
			ixx := append(append([]int{}, ix...), i)
			swalk(fld.Type, ixx, indent+". ")
		}
	default:
		// everything is either a struct or a field (that's an assumption about Options)
		fields = append(fields, field{ix})
	}
}

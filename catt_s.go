package ctoken

// This file implements interface Stringser.

import (
	"encoding/xml"
	_ "github.com/fbaube/stringutils"
)

// Alias the standard library's XML type
// (for simplicity and convenience) to
//   - attach methods to it (e.g. interface
//     [stringutils.Stringser]), and
//   - use it for other markups too (like Markdown)
//
// type xml.Attr struct { Name xml.Name; Value string }
// .
type CAtt xml.Attr
type CAtts []CAtt

func (A CAtt) Echo() string {
	return " " + CName(A.Name).Echo() + "=\"" + A.Value + "\""
}

func (AL CAtts) Echo() string {
	var s string
	for _, A := range AL {
		s += " " + CName(A.Name).Echo() + "=\"" + A.Value + "\""
	}
	return s
}

func (A CAtt) Info() string {
	return A.Echo()
}

func (A CAtt) Debug() string {
	return A.Info()
}

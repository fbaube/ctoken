package ctoken

import (
	"encoding/xml"
	SU "github.com/fbaube/stringutils"
	S "strings"
)

// Alias the standard library's XML type
// (for simplicity and convenience) to
//   - attach methods to it (e.g. interface [SU.Stringser]), and
//   - use it for other markups too (like Markdown)
type CName xml.Name

// CName is an [xml.Name].
//
// type xml.Name struct { Space, Local string }
// .
func (p1 *CName) Equals(p2 *CName) bool {
	return p1.Space == p2.Space && p1.Local == p2.Local
}

func (p *CName) FixNS() {
	if p.Space != "" && !S.HasSuffix(p.Space, ":") {
		p.Space = p.Space + SU.Trim(":")
	}
}

// NewCName enforcs a colon after a non-empty
// namespace if it is not there already.
func NewCName(ns, local string) *CName {
	p := new(CName)
	if ns != "" && !S.HasSuffix(ns, ":") {
		ns += ":"
	}
	p.Space = ns
	p.Local = local
	return p
}

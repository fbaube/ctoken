package ctoken

// This file implements interface Stringser.

import (
       "time"
	"encoding/xml"
	SU "github.com/fbaube/stringutils"
	"github.com/fbaube/miscutils"
)

var ttt time.Time

// init uses [miscutils.Into] and [Into] 
func init() {
      // func Into(s string) time.Time
      ttt = miscutils.Into("IN")
}

// Alias the standard library's XML type
// (for simplicity and convenience) to
//   - attach methods to it (e.g. interface
//     [github.com/fbaube/stringutils.Stringser]), and
//     [stringutils.Stringser]), and
//     [SU.Stringser]), and
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
		s += " " + CName(A.Name).Echo() + "=\"" + A.Value + "\"" +
		SU.Yn(true)
	}
	return s
}

func (A CAtt) Info() string {
	return A.Echo()
}

func (A CAtt) Debug() string {
	return A.Info()
}
func (x1 CAtts) AsStdLibXml() []xml.Attr {
	var x2 []CAtt
	var x3 []xml.Attr
	x2 = x1
	// x3 = []xml.Attr(x2)
	for _, A := range x2 {
		x3 = append(x3, xml.Attr(A))
	}
	return x3
}

package ctoken

import (
	"github.com/nbio/xml"
	"fmt"
)

// FilePosition is a char.position (e.g. from xml.Decoder)
// plus line nr & column nr (when they can be calculated).
//
// FilePosition implements interface [SU.Stringser].
// .
type FilePosition struct {
	// Pos is the byte position in file,
	// e.g. from xml.Decoder.InputOffset()
	Pos int
	// Lnr & Col are line nr & column nr
	Lnr, Col int
}

// NewFilePosition takes & uses only
// the character position in the file.
func NewFilePosition(i int) *FilePosition {
	p := new(FilePosition)
	p.Pos = i
	return p
}

func (fp FilePosition) Info() string {
	if fp.Lnr == 0 && fp.Col == 0 {
		return fmt.Sprintf("ch%03d", fp.Pos)
	}
	return fmt.Sprintf("ch%d:L%02dc%02d", fp.Pos, fp.Lnr, fp.Col)
}

func (fp FilePosition) Echo() string {
     	// A Bogus call to let us include SU
	return fp.Info()
}

func (fp FilePosition) Debug() string {
	return fmt.Sprintf("fileposn:ch%d,L%02d.c%02d",
		fp.Pos, fp.Lnr, fp.Col)
}

type FileRange struct {
	Beg FilePosition
	End FilePosition
}

// === SPAN ===

// Span specifies the range of a subset of a string (that is not included in the struct).
//
// Span implements interface [SU.Stringser].
//
// FIXME:  Make this a ptr to a ContentityNode
// .
type Span struct {
	TagName string
	Atts    []xml.Attr
	// SliceBounds
	FileRange
}

func (sp Span) GetSpanOfString(s string) string {
	if len(s) == 0 {
		panic("Zero-len Raw")
	}
	if sp.End.Pos == 0 {
		return ""
	}
	if sp.End.Pos == -1 && sp.Beg.Pos == 0 {
		return s
	}
	if sp.End.Pos > len(s) {
		panic(fmt.Sprintf("Span: END %d > LEN %d",
			sp.End.Pos, len(s)))
	}
	if sp.Beg.Pos > sp.End.Pos {
		panic(fmt.Sprintf("Span: BEG %d > END %d",
			sp.Beg.Pos, sp.End.Pos))
	}
	return s[sp.Beg.Pos:sp.End.Pos]
}

func (sp Span) Info() string {
	return fmt.Sprintf("%s[%d:%d]", sp.TagName, sp.Beg.Pos, sp.End.Pos)
}

func (sp Span) Echo() string {
	return sp.Echo()
}

func (sp Span) Debug() string {
	return fmt.Sprintf("<%s>[%s:%s]",
		sp.TagName, sp.Beg.Debug(), sp.End.Debug())
}

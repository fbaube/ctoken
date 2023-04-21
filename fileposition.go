package ctoken

import (
	"fmt"
	_ "github.com/fbaube/stringutils"
)

// FilePosition is a char.position (e.g. from xml.Decoder)
// plus line nr & column nr (when they can be calculated).
//
// FilePosition implements interface [stringutils.Stringser].
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
	return fp.Info()
}

func (fp FilePosition) Debug() string {
	return fmt.Sprintf("fileposn:ch%d,L%02d.c%02d",
		fp.Pos, fp.Lnr, fp.Col)
}

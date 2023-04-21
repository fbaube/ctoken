package ctoken

import (
	// "encoding/xml"
	"fmt"
)

// LAToken is a location-aware XML token.
type LAToken struct {
	CToken
	FilePosition
}

func (fp FilePosition) String() string {
	if fp.Lnr == 0 && fp.Col == 0 {
		return fmt.Sprintf("[%03d]", fp.Pos)
	}
	return fmt.Sprintf("[%d](L%02dc%02d)", fp.Pos, fp.Lnr, fp.Col)
}

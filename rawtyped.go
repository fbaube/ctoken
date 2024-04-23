package ctoken

// This should file should be named
// typedraw.go not rawtyped.go but
// that gets redd wrong, as "type draw".

import SU "github.com/fbaube/stringutils"

type Raw string

// TypedRaw includes [stringutils.MarkupType] 
// and can have it set to [MU_type_DIRLIKE].
type TypedRaw struct {
	Raw
	// We have to rename this field so that 
	// we don't confuse the compiler too much.
	RawMT SU.MarkupType
}

func (s Raw) S() string {
	return string(s)
}

func (p *TypedRaw) S() string {
	return string(p.Raw)
}

// RawType is a convenience function so that
// if (i.e. when) it becomes convenient, the
// elements of [TypedRaw] can be unexported.
// .
func (p *TypedRaw) RawType() SU.MarkupType {
	return p.RawMT
}

// IsDirlike is IsDir()-like but more general. Dirlike 
// is shorthand for "cannot (is not allowed to!) have 
// own content", but it can be defined as "is/has link(s) 
// to other stuff" - i.e. a directory or a symbolic link.
// In this context (i.e. when embedded in TypedRaw), it
// means SU.MU_type_DIRLIKE
// .
func (p *TypedRaw) IsDirlike() bool {
     return p.RawMT == SU.MU_type_DIRLIKE
}
package ctoken

import SU "github.com/fbaube/stringutils"

type Raw string

// TypedRaw includes [stringutils.MarkupType] 
// and can have it set to [MU_type_DIRLIKE].
type TypedRaw struct {
	Raw
	SU.MarkupType
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
	return p.MarkupType
}

// IsDirlike is like IsDir() but more general. Dirlike 
// is shorthand for "cannot (is not allowed to!) have 
// own content", but it can be defined as "is/has link(s) 
// to other stuff" - i.e. a directory or a symbolic link.
// In this context (i.e. when embedded in TypedRaw), it
// means SU.MU_type_DIRLIKE
// .
func (p *TypedRaw) IsDirlike() bool {
     return p.MarkupType == SU.MU_type_DIRLIKE
}
package ctoken

import (
	// "io"
	S "strings"
)

func (N CName) Echo() string {
	// Enforce colon at the end of `N.Space`
	if N.Space != "" && !S.HasSuffix(N.Space, ":") {
		return N.Space + ":" + N.Local
	}
	return N.Space + N.Local
}

func (N CName) Info() string {
	return N.Echo()
}

func (N CName) Debug() string {
	return N.Info()
}

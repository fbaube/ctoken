package ctoken

// import "encoding/xml"

// GetFirstCTokenByTag checks a start-element's tag's local
// name only, not any namespace. If no match, it returns nil.
// This func returns only a naked CToken, taken from a slice
// and probably without context, so it is meant only for
// processing XML catalog files. General XML processing
// should use the GToken version, which returns a CToken
// in the context of a tree structure.
// .
func GetFirstCTokenByTag(tkzn []CToken, s string) *CToken {
	if s == "" {
		return nil
	}
	for _, CT := range tkzn {
		if CT.TDType == TD_type_ELMNT && CT.CName.Local == s {
			if CT.Text != s {
				panic("Mismatch in GetFirstCTokenByTag: " + s)
			}
			return &(CT)
		}
	}
	return nil
}

// GetAllCTokensByTag checks the basic tag only, not any namespace.
// .
func GetAllCTokensByTag(tkzn []CToken, s string) []CToken {
	if s == "" {
		return nil
	}
	var ret []CToken
	ret = make([]CToken, 0)
	for _, p := range tkzn {
		// if SE, OK := p.(xml.StartElement); OK {
		if p.TDType == TD_type_ELMNT {
			if p.CName.Local == s {
				// fmt.Printf("found a match [%d] %s (NS:%s)\n", i, p.GName.Local, p.GName.Space)
				ret = append(ret, p)
			}
		}
	}
	return ret
}

/* OBS ?

// GetAttVal returns the attribute's string value, or "" if not found.
func GetAttVal(se xml.StartElement, att string) string {
	for _, A := range se.Attr {
		if A.Name.Local == att {
			return A.Value
		}
	}
	return ""
}

// GetXAttVal returns the attribute's string value, or "" if not found.
func (ct CToken) GetXAttVal(att string) string {
	for _, A := range ct.CAtts {
		if A.Name.Local == att {
			return A.Value
		}
	}
	return ""
}

*/

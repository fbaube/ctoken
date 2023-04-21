package ctoken

import ()

func (ct CToken) Echo() string {
	return "(ctkn.echo)"
}

func (ct CToken) Info() string {
	var s string
	s = ct.TDType.S() + ":"
	switch ct.TDType {
	case TD_type_ELMNT, TD_type_ENDLM:
		s += ct.CName.Info()
	case TD_type_CDATA, TD_type_COMNT:
		s += "\"" + ct.Text + "\""
	case TD_type_PINST:
		s += ct.ControlStrings[0] + ":" + "<" + ct.Text + ">"
	case TD_type_DRCTV:
		s += ct.ControlStrings[0] + ":" +
			ct.ControlStrings[1] + ":" +
			"<" + ct.Text + ">"
	default:
		panic("ctoken.Info()!")
	}
	return s
}

func (ct CToken) Debug() string {
	return ct.Info()
}

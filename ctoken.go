package ctoken

import (
	"encoding/xml"
	"fmt"
	SU "github.com/fbaube/stringutils"
	// XU "github.com/fbaube/xmlutils"
	L "github.com/fbaube/mlog"
	S "strings"
)

// CToken is the lowest common denominator of tokens parsed
// from XML mixed content and other content-oriented markup.
// It has [stringutils.MarkupType].
//
// CToken:
//   - Common Token
//   - Content Token
//   - Combined Token
//   - Canonical Token
//   - Consolidated Token
//   - ConMuchoGusto Token :-P
//
// A CToken contains all that can be parsed from a token that
// is considered in isolation, as-is, without the context of
// surrounding markup. It should record/reflect/reproduce any
// XML (or HTML) token faithfully, and also accommodate any
// token from Markdown or (in the future) related markup
// such as Docbook or Asciidoc or RST (restructured text).
//
// The use of an XML-like data structure as the lingua franca
// is also meant to make XML-style automated processing simpler.
//
// The use of a single unified token representation is intended
// most of all to simplify & unify tokenisation across LwDITA's
// three supported input formats: XDITA XML, HDITA HTML5, and
// MDITA-XP Markdown. It also serves to represent all the
// various kinds of XML directives, including DTDs(!).
//
// Creation of a new CToken from an [encoding/xml.Token] is
// by design very straightforward, but creation from other
// types of token, such as HTML or Markdown, must be done
// in their other packages in order to prevent circular
// dependencies.
//
// For convenience & simplicity, some items in the struct
// are simply aliases for Go's XML structs, but then these
// must also be adaptable for Markdown. For example, when
// Pandoc-style attributes are used.
//
// CToken implements interface [stringutils.Stringser].
// .
type CToken struct {
	// ==================================
	// The original ("source code") token,
	// and other information about it
	// ==================================
	// SourceToken is the original token.
	// Keep it around "just in case".
	// TODO: Make this an Echoer !
	// Types:
	//  - XML: [xml.Token] from [xml.Decoder]
	//  - HTML: TBS
	//  - Markdown: TBS
	// Note that an XML Token is transitory,
	// so every Token has to be cloned, by
	// calling [xml.CopyToken].
	SourceToken interface{}
	// MarkupType of the original token; the value is
	// one of MU_type_(XML/HTML/MKDN/BIN/SQL/DIRLIKE). 
	// It is particularly helpful to have this info at the
	// token level when we consider that for example, we can
	// embed HTML tags in Markdown. Note that in the future,
	// each value could actually be a namespace declaration.
	SU.MarkupType
	// FilePosition is char position, and line nr & column nr.
	FilePosition

	// TDType comprises (a) the types of [xml.Token]
	// (they are all different struct's, actually),
	// plus (b) the (sub)types of [xml.Directive].
	// Note that [TD_type_ENDLM] ("EndElement") is
	// superfluous when token depth info is available.
	TDType
	// CName is ONLY for elements
	// (i.e. [TD_type_ELMNT] and [TD_type_ENDLM]).
	CName
	// CAtts is ONLY for [TD_type_ELMNT].
	CAtts
	// Text holds CDATA, and a PI's Instruction,
	// and a DOCTYPE's root element declaration,
	// and
	Text string
	// ControlStrings is tipicly XML PI & Directive stuff.
	// When it is used, its length is 1 or 2.
	//  - XML PI: the Target field
	//  - XML directive: the directive subtype
	// But this field also available for other data that
	// is not classifiable as source text.
	ControlStrings []string
}

/* REF
type GToken struct {
	// Depth is the level of nesting of the source tag.
	Depth int
	// IsBlock and IsInline are
	// dupes of TagalogEntry ?
	IsBlock, IsInline bool
        // NodeLevel is dupe of Depth ?
	NodeLevel int
        // Lotsa info ?
	*lwdx.TagalogEntry
	// DitaTag and HtmlTag are
	// dupes of TagalogEntry ?
	NodeKind, DitaTag, HtmlTag, NodeText string
}
*/

// NewCTokenFromXmlToken returns a single token type that replaces
// the unwieldy multi-typed mess of the standard library.
//
// It returns a nil ptr for an ignorable, skippable token, like
// all-whitespace.
// .
func NewCTokenFromXmlToken(XT xml.Token) *CToken {
	ctkn := new(CToken)
	ctkn.SourceToken = XT
	switch XT.(type) {
	case xml.StartElement:
		ctkn.TDType = TD_type_ELMNT
		// type xml.StartElement struct {
		//     Name Name ; Attr []Attr }
		var xSE xml.StartElement
		xSE = xml.CopyToken(XT).(xml.StartElement)
		// ctkn.TagOrDirective = xSE.Name.Local
		// ctkn.Strings = []string{xSE.Name.Local}
		ctkn.Text = xSE.Name.Local
		ctkn.CName = CName(xSE.Name)
		ctkn.CName.FixNS()
		// println("Elm:", ctkn.CName.String())

		// TODO: Is this the place check for any of the other
		// "standard" XML namespaces that we might encounter ?
		if ctkn.CName.Space == NS_XML {
			ctkn.CName.Space = "xml:"
		}
		for _, xA := range xSE.Attr {
			if xA.Name.Space == NS_XML {
				// println("TODO check name.local:
				// newgtoken xml:" + A.Name.Local)
				xA.Name.Space = "xml:"
			}
			gA := CAtt(xA)
			ctkn.CAtts = append(ctkn.CAtts, gA)
		}

	case xml.EndElement:
		// An EndElement has a Name (CName).
		ctkn.TDType = TD_type_ENDLM
		// type xml.EndElement struct { Name Name }
		var xEE xml.EndElement
		xEE = xml.CopyToken(XT).(xml.EndElement)
		// ctkn.TagOrDirective = xEE.Name.Local
		// ctkn.Strings = []string{xEE.Name.Local}
		ctkn.Text = xEE.Name.Local
		ctkn.CName = CName(xEE.Name)
		if ctkn.CName.Space == NS_XML {
			ctkn.CName.Space = "xml:"
		}
		// fmt.Printf("<!--End-Tagnt--> %s \n", outGT.Echo())

	case xml.Comment:
		// type Comment []byte
		ctkn.TDType = TD_type_COMNT
		// ctkn.TagOrDirective = "//" // TD_type_COMNT
		// ctkn.DirectiveText = S.TrimSpace(
		//	string([]byte(XT.(xml.Comment))))
		// ctkn.Strings = []string{S.TrimSpace(
		ctkn.Text = S.TrimSpace(string([]byte(XT.(xml.Comment))))
		// fmt.Printf("<!-- Comment --> <!-- %s --> \n",
		//     outGT.DirectiveText)

	case xml.ProcInst:
		ctkn.TDType = TD_type_PINST
		// type xml.ProcInst struct { Target string ; Inst []byte }
		xPI := XT.(xml.ProcInst)
		// ctkn.TagOrDirective = S.TrimSpace(xPI.Target)
		// ctkn.DirectiveText = S.TrimSpace(string(xPI.Inst))
		// prinst := S.TrimSpace(string(xPI.Inst))
		// ctkn.Strings = []string{target, prinst}
		ctkn.Text = S.TrimSpace(string(xPI.Inst))
		target := S.TrimSpace(xPI.Target)
		ctkn.ControlStrings = []string{target}
		// 2023.04 This works OK :-D
		if target == "xml" {
			// fmt.Printf("XML!! %s \n", ctkn.DirectiveText)
		}
		// fmt.Printf("<!--ProcInstr--> <?%s %s?> \n",
		// 	outGT.Keyword, outGT.DirectiveText)

	case xml.Directive: // type Directive []byte
		ctkn.TDType = TD_type_DRCTV
		var fullDrctv, string0, tmp string
		fullDrctv = S.TrimSpace(string([]byte(XT.(xml.Directive))))
		// ctkn.Strings = make([]string, 3)
		string0, tmp = SU.SplitOffFirstWord(fullDrctv)
		// TODO: Assign TagOrDirective to ctkn.TDType
		// 2023.04 This works OK :-D
		if string0 != "DOCTYPE" {
			ctkn.ControlStrings = make([]string, 1)
			ctkn.ControlStrings[0] = string0
			ctkn.Text = tmp
			fmt.Printf("newCtkn L212 (!Drctv): %s: %s||%s \n",
				ctkn.TDType, string0, tmp)
		} else {
			ctkn.ControlStrings = make([]string, 2)
			ctkn.ControlStrings[0] = string0
			ctkn.ControlStrings[1], ctkn.Text =
				SU.SplitOffFirstWord(tmp)
			L.L.Okay("NewCtoken: Directive: %s: %s||%s||%s",
				ctkn.TDType, string0,
				ctkn.ControlStrings[1], tmp)
		}
		// fmt.Printf("NewCtkn: Drctv: [0] %s [1] %s [2] %s",
		//	ctkn.Strings[0], ctkn.Strings[1], ctkn.Strings[2])
		// fmt.Printf("<!--Directive--> <!%s %s> \n",
		// 	outGT.Keyword, outGT.Otherwo rds)

	case xml.CharData:
		// type CharData []byte
		ctkn.TDType = TD_type_CDATA
		bb := []byte(xml.CopyToken(XT).(xml.CharData))
		ss := S.TrimSpace(string(bb))
		// This might cause problems in a scenario
		// where we actually have to worry about
		// the finer points of whitespace handing.
		// But ignore it for now, to preserve sanity.
		if ss == "" {
			return nil
		}
		ctkn.Text = ss
		// fmt.Printf("<!--Char-Data--> %s \n", outGT.DirectiveText)

	default:
		ctkn.TDType = TD_type_ERROR
		// L.L.Error("Unrecognized xml.Token type<%T> for: %+v", XT, XT)
		// continue
	}
	return ctkn
}

func (ct CToken) IsNonElement() bool {
	switch ct.TDType {
	case TD_type_DOCMT, TD_type_ELMNT, TD_type_ENDLM, TD_type_VOIDD:
		return false
	case TD_type_CDATA, TD_type_PINST, TD_type_COMNT, TD_type_DRCTV,
		// DIRECTIVE SUBTYPES
		TD_type_Doctype, TD_type_Element, TD_type_Attlist,
		TD_type_Entity, TD_type_Notation,
		// TBD/experimental
		TD_type_ID, TD_type_IDREF, TD_type_Enum:
		return true
	case TD_type_ERROR:
		panic("XU.IsNonElement")
	}
	return true
}

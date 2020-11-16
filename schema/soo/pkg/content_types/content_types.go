//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package content_types

import (
	_f "encoding/xml"
	_bd "fmt"
	_be "github.com/unidoc/unioffice"
	_c "regexp"
)

func NewCT_Types() *CT_Types { _bfa := &CT_Types{}; return _bfa }

// ValidateWithPath validates the Default and its children, prefixing error messages with path
func (_ffec *Default) ValidateWithPath(path string) error {
	if _gab := _ffec.CT_Default.ValidateWithPath(path); _gab != nil {
		return _gab
	}
	return nil
}

func (_gc *Override) MarshalXML(e *_f.Encoder, start _f.StartElement) error {
	return _gc.CT_Override.MarshalXML(e, start)
}

// Validate validates the CT_Types and its children
func (_bb *CT_Types) Validate() error {
	return _bb.ValidateWithPath("\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073")
}

func NewTypes() *Types { _gcd := &Types{}; _gcd.CT_Types = *NewCT_Types(); return _gcd }

func (_ccf *CT_Types) MarshalXML(e *_f.Encoder, start _f.StartElement) error {
	e.EncodeToken(start)
	if _ccf.Default != nil {
		_dc := _f.StartElement{Name: _f.Name{Local: "\u0044e\u0066\u0061\u0075\u006c\u0074"}}
		for _, _gd := range _ccf.Default {
			e.EncodeElement(_gd, _dc)
		}
	}
	if _ccf.Override != nil {
		_ef := _f.StartElement{Name: _f.Name{Local: "\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}}
		for _, _cca := range _ccf.Override {
			e.EncodeElement(_cca, _ef)
		}
	}
	e.EncodeToken(_f.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the Override and its children, prefixing error messages with path
func (_cba *Override) ValidateWithPath(path string) error {
	if _cfac := _cba.CT_Override.ValidateWithPath(path); _cfac != nil {
		return _cfac
	}
	return nil
}

func (_bad *Types) UnmarshalXML(d *_f.Decoder, start _f.StartElement) error {
	_bad.CT_Types = *NewCT_Types()
_gbb:
	for {
		_bae, _cdb := d.Token()
		if _cdb != nil {
			return _cdb
		}
		switch _abe := _bae.(type) {
		case _f.StartElement:
			switch _abe.Name {
			case _f.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", Local: "\u0044e\u0066\u0061\u0075\u006c\u0074"}:
				_ffea := NewDefault()
				if _ddeg := d.DecodeElement(_ffea, &_abe); _ddeg != nil {
					return _ddeg
				}
				_bad.Default = append(_bad.Default, _ffea)
			case _f.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", Local: "\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:
				_fee := NewOverride()
				if _gcdg := d.DecodeElement(_fee, &_abe); _gcdg != nil {
					return _gcdg
				}
				_bad.Override = append(_bad.Override, _fee)
			default:
				_be.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u0054\u0079\u0070e\u0073 \u0025\u0076", _abe.Name)
				if _acf := d.Skip(); _acf != nil {
					return _acf
				}
			}
		case _f.EndElement:
			break _gbb
		case _f.CharData:
		}
	}
	return nil
}

func NewCT_Default() *CT_Default {
	_a := &CT_Default{}
	_a.ExtensionAttr = "\u0078\u006d\u006c"
	_a.ContentTypeAttr = "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c"
	return _a
}

// ValidateWithPath validates the CT_Types and its children, prefixing error messages with path
func (_cbc *CT_Types) ValidateWithPath(path string) error {
	for _fg, _cg := range _cbc.Default {
		if _fe := _cg.ValidateWithPath(_bd.Sprintf("\u0025\u0073\u002f\u0044\u0065\u0066\u0061\u0075\u006ct\u005b\u0025\u0064\u005d", path, _fg)); _fe != nil {
			return _fe
		}
	}
	for _gdf, _cfa := range _cbc.Override {
		if _eeb := _cfa.ValidateWithPath(_bd.Sprintf("\u0025s\u002fO\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u005b\u0025\u0064\u005d", path, _gdf)); _eeb != nil {
			return _eeb
		}
	}
	return nil
}

type Default struct{ CT_Default }

// Validate validates the Default and its children
func (_eff *Default) Validate() error {
	return _eff.ValidateWithPath("\u0044e\u0066\u0061\u0075\u006c\u0074")
}

var ST_ContentTypePatternRe = _c.MustCompile(ST_ContentTypePattern)

func (_bed *CT_Default) UnmarshalXML(d *_f.Decoder, start _f.StartElement) error {
	_bed.ExtensionAttr = "\u0078\u006d\u006c"
	_bed.ContentTypeAttr = "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c"
	for _, _cd := range start.Attr {
		if _cd.Name.Local == "\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn" {
			_cdf, _e := _cd.Value, error(nil)
			if _e != nil {
				return _e
			}
			_bed.ExtensionAttr = _cdf
			continue
		}
		if _cd.Name.Local == "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065" {
			_ff, _ee := _cd.Value, error(nil)
			if _ee != nil {
				return _ee
			}
			_bed.ContentTypeAttr = _ff
			continue
		}
	}
	for {
		_bf, _ec := d.Token()
		if _ec != nil {
			return _bd.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020%\u0073", _ec)
		}
		if _ffe, _cdfd := _bf.(_f.EndElement); _cdfd && _ffe.Name == start.Name {
			break
		}
	}
	return nil
}

func (_bedd *CT_Types) UnmarshalXML(d *_f.Decoder, start _f.StartElement) error {
_cb:
	for {
		_dea, _fcf := d.Token()
		if _fcf != nil {
			return _fcf
		}
		switch _bff := _dea.(type) {
		case _f.StartElement:
			switch _bff.Name {
			case _f.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", Local: "\u0044e\u0066\u0061\u0075\u006c\u0074"}:
				_ga := NewDefault()
				if _abf := d.DecodeElement(_ga, &_bff); _abf != nil {
					return _abf
				}
				_bedd.Default = append(_bedd.Default, _ga)
			case _f.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", Local: "\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:
				_da := NewOverride()
				if _ca := d.DecodeElement(_da, &_bff); _ca != nil {
					return _ca
				}
				_bedd.Override = append(_bedd.Override, _da)
			default:
				_be.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073\u0020\u0025\u0076", _bff.Name)
				if _dg := d.Skip(); _dg != nil {
					return _dg
				}
			}
		case _f.EndElement:
			break _cb
		case _f.CharData:
		}
	}
	return nil
}

type Types struct{ CT_Types }

func NewOverride() *Override { _ged := &Override{}; _ged.CT_Override = *NewCT_Override(); return _ged }

const ST_ExtensionPattern = "\u0028\u005b\u0021\u0024\u0026\u0027\\\u0028\u005c\u0029\u005c\u002a\\\u002b\u002c\u003a\u003d\u005d\u007c(\u0025\u005b\u0030\u002d\u0039a\u002d\u0066\u0041\u002d\u0046\u005d\u005b\u0030\u002d\u0039\u0061\u002df\u0041\u002d\u0046\u005d\u0029\u007c\u005b\u003a\u0040\u005d\u007c\u005b\u0061\u002d\u007aA\u002d\u005a\u0030\u002d\u0039\u005c\u002d\u005f~\u005d\u0029\u002b"

func (_ce *Default) MarshalXML(e *_f.Encoder, start _f.StartElement) error {
	return _ce.CT_Default.MarshalXML(e, start)
}

type CT_Types struct {
	Default  []*Default
	Override []*Override
}

// Validate validates the Override and its children
func (_ece *Override) Validate() error {
	return _ece.ValidateWithPath("\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065")
}

type CT_Default struct {
	ExtensionAttr   string
	ContentTypeAttr string
}

var ST_ExtensionPatternRe = _c.MustCompile(ST_ExtensionPattern)

func NewDefault() *Default { _gb := &Default{}; _gb.CT_Default = *NewCT_Default(); return _gb }

func (_cdfa *Types) MarshalXML(e *_f.Encoder, start _f.StartElement) error {
	start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s"})
	start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0054\u0079\u0070e\u0073"
	return _cdfa.CT_Types.MarshalXML(e, start)
}

type CT_Override struct {
	ContentTypeAttr string
	PartNameAttr    string
}

func NewCT_Override() *CT_Override {
	_de := &CT_Override{}
	_de.ContentTypeAttr = "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c"
	return _de
}

// Validate validates the CT_Default and its children
func (_d *CT_Default) Validate() error {
	return _d.ValidateWithPath("\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074")
}

const ST_ContentTypePattern = "\u005e\\\u0070{\u004c\u0061\u0074\u0069\u006e\u007d\u002b\u002f\u002e\u002a\u0024"

type Override struct{ CT_Override }

// ValidateWithPath validates the Types and its children, prefixing error messages with path
func (_agg *Types) ValidateWithPath(path string) error {
	if _efa := _agg.CT_Types.ValidateWithPath(path); _efa != nil {
		return _efa
	}
	return nil
}

func (_bee *CT_Override) UnmarshalXML(d *_f.Decoder, start _f.StartElement) error {
	_bee.ContentTypeAttr = "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c"
	for _, _ed := range start.Attr {
		if _ed.Name.Local == "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065" {
			_bea, _eg := _ed.Value, error(nil)
			if _eg != nil {
				return _eg
			}
			_bee.ContentTypeAttr = _bea
			continue
		}
		if _ed.Name.Local == "\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065" {
			_af, _dd := _ed.Value, error(nil)
			if _dd != nil {
				return _dd
			}
			_bee.PartNameAttr = _af
			continue
		}
	}
	for {
		_aba, _db := d.Token()
		if _db != nil {
			return _bd.Errorf("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u003a\u0020\u0025\u0073", _db)
		}
		if _ffc, _cc := _aba.(_f.EndElement); _cc && _ffc.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Default and its children, prefixing error messages with path
func (_ab *CT_Default) ValidateWithPath(path string) error {
	if !ST_ExtensionPatternRe.MatchString(_ab.ExtensionAttr) {
		return _bd.Errorf("\u0025s\u002f\u006d.\u0045\u0078\u0074\u0065n\u0073\u0069\u006fn\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074 m\u0061\u0074\u0063h\u0020\u0027%\u0073\u0027\u0020\u0028\u0068\u0061v\u0065\u0020%\u0076\u0029", path, ST_ExtensionPatternRe, _ab.ExtensionAttr)
	}
	if !ST_ContentTypePatternRe.MatchString(_ab.ContentTypeAttr) {
		return _bd.Errorf("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, ST_ContentTypePatternRe, _ab.ContentTypeAttr)
	}
	return nil
}

// Validate validates the CT_Override and its children
func (_fc *CT_Override) Validate() error {
	return _fc.ValidateWithPath("C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065")
}

// ValidateWithPath validates the CT_Override and its children, prefixing error messages with path
func (_bg *CT_Override) ValidateWithPath(path string) error {
	if !ST_ContentTypePatternRe.MatchString(_bg.ContentTypeAttr) {
		return _bd.Errorf("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, ST_ContentTypePatternRe, _bg.ContentTypeAttr)
	}
	return nil
}

func (_ge *CT_Override) MarshalXML(e *_f.Encoder, start _f.StartElement) error {
	start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"}, Value: _bd.Sprintf("\u0025\u0076", _ge.ContentTypeAttr)})
	start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"}, Value: _bd.Sprintf("\u0025\u0076", _ge.PartNameAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_f.EndElement{Name: start.Name})
	return nil
}

func (_g *CT_Default) MarshalXML(e *_f.Encoder, start _f.StartElement) error {
	start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"}, Value: _bd.Sprintf("\u0025\u0076", _g.ExtensionAttr)})
	start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"}, Value: _bd.Sprintf("\u0025\u0076", _g.ContentTypeAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_f.EndElement{Name: start.Name})
	return nil
}

func (_bc *Default) UnmarshalXML(d *_f.Decoder, start _f.StartElement) error {
	_bc.CT_Default = *NewCT_Default()
	for _, _dde := range start.Attr {
		if _dde.Name.Local == "\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn" {
			_aa, _ae := _dde.Value, error(nil)
			if _ae != nil {
				return _ae
			}
			_bc.ExtensionAttr = _aa
			continue
		}
		if _dde.Name.Local == "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065" {
			_aec, _ecd := _dde.Value, error(nil)
			if _ecd != nil {
				return _ecd
			}
			_bc.ContentTypeAttr = _aec
			continue
		}
	}
	for {
		_gaf, _fd := d.Token()
		if _fd != nil {
			return _bd.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020\u0025\u0073", _fd)
		}
		if _dba, _dbc := _gaf.(_f.EndElement); _dbc && _dba.Name == start.Name {
			break
		}
	}
	return nil
}

func (_bbd *Override) UnmarshalXML(d *_f.Decoder, start _f.StartElement) error {
	_bbd.CT_Override = *NewCT_Override()
	for _, _ffd := range start.Attr {
		if _ffd.Name.Local == "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065" {
			_dbd, _ecg := _ffd.Value, error(nil)
			if _ecg != nil {
				return _ecg
			}
			_bbd.ContentTypeAttr = _dbd
			continue
		}
		if _ffd.Name.Local == "\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065" {
			_cbg, _ba := _ffd.Value, error(nil)
			if _ba != nil {
				return _ba
			}
			_bbd.PartNameAttr = _cbg
			continue
		}
	}
	for {
		_fef, _abfa := d.Token()
		if _abfa != nil {
			return _bd.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u004f\u0076\u0065r\u0072\u0069\u0064\u0065: \u0025\u0073", _abfa)
		}
		if _gg, _cgf := _fef.(_f.EndElement); _cgf && _gg.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Types and its children
func (_eda *Types) Validate() error { return _eda.ValidateWithPath("\u0054\u0079\u0070e\u0073") }

func init() {
	_be.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073", NewCT_Types)
	_be.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074", NewCT_Default)
	_be.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065", NewCT_Override)
	_be.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "\u0054\u0079\u0070e\u0073", NewTypes)
	_be.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "\u0044e\u0066\u0061\u0075\u006c\u0074", NewDefault)
	_be.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065", NewOverride)
}

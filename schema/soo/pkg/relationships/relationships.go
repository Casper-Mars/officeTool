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

package relationships

import (
	_ga "encoding/xml"
	_d "fmt"
	_gb "github.com/unidoc/unioffice"
)

type CT_Relationship struct {
	TargetModeAttr ST_TargetMode
	TargetAttr     string
	TypeAttr       string
	IdAttr         string
	Content        string
}

func NewCT_Relationships() *CT_Relationships { _dc := &CT_Relationships{}; return _dc }

func (_f *CT_Relationship) MarshalXML(e *_ga.Encoder, start _ga.StartElement) error {
	if _f.TargetModeAttr != ST_TargetModeUnset {
		_c, _b := _f.TargetModeAttr.MarshalXMLAttr(_ga.Name{Local: "\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065"})
		if _b != nil {
			return _b
		}
		start.Attr = append(start.Attr, _c)
	}
	start.Attr = append(start.Attr, _ga.Attr{Name: _ga.Name{Local: "\u0054\u0061\u0072\u0067\u0065\u0074"}, Value: _d.Sprintf("\u0025\u0076", _f.TargetAttr)})
	start.Attr = append(start.Attr, _ga.Attr{Name: _ga.Name{Local: "\u0054\u0079\u0070\u0065"}, Value: _d.Sprintf("\u0025\u0076", _f.TypeAttr)})
	start.Attr = append(start.Attr, _ga.Attr{Name: _ga.Name{Local: "\u0049\u0064"}, Value: _d.Sprintf("\u0025\u0076", _f.IdAttr)})
	e.EncodeElement(_f.Content, start)
	e.EncodeToken(_ga.EndElement{Name: start.Name})
	return nil
}

func NewRelationships() *Relationships {
	_ebc := &Relationships{}
	_ebc.CT_Relationships = *NewCT_Relationships()
	return _ebc
}

func (_dfc *Relationship) UnmarshalXML(d *_ga.Decoder, start _ga.StartElement) error {
	_dfc.CT_Relationship = *NewCT_Relationship()
	for _, _dcg := range start.Attr {
		if _dcg.Name.Local == "\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065" {
			_dfc.TargetModeAttr.UnmarshalXMLAttr(_dcg)
			continue
		}
		if _dcg.Name.Local == "\u0054\u0061\u0072\u0067\u0065\u0074" {
			_aag, _cf := _dcg.Value, error(nil)
			if _cf != nil {
				return _cf
			}
			_dfc.TargetAttr = _aag
			continue
		}
		if _dcg.Name.Local == "\u0054\u0079\u0070\u0065" {
			_cee, _ca := _dcg.Value, error(nil)
			if _ca != nil {
				return _ca
			}
			_dfc.TypeAttr = _cee
			continue
		}
		if _dcg.Name.Local == "\u0049\u0064" {
			_ec, _af := _dcg.Value, error(nil)
			if _af != nil {
				return _af
			}
			_dfc.IdAttr = _ec
			continue
		}
	}
	for {
		_gd, _fe := d.Token()
		if _fe != nil {
			return _d.Errorf("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0052\u0065\u006c\u0061\u0074i\u006f\u006e\u0073\u0068\u0069\u0070\u003a\u0020\u0025\u0073", _fe)
		}
		if _gag, _feb := _gd.(_ga.EndElement); _feb && _gag.Name == start.Name {
			break
		}
	}
	return nil
}

type ST_TargetMode byte

func (_cbdf *ST_TargetMode) UnmarshalXML(d *_ga.Decoder, start _ga.StartElement) error {
	_gg, _cg := d.Token()
	if _cg != nil {
		return _cg
	}
	if _eaf, _eab := _gg.(_ga.EndElement); _eab && _eaf.Name == start.Name {
		*_cbdf = 1
		return nil
	}
	if _ab, _aae := _gg.(_ga.CharData); !_aae {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gg)
	} else {
		switch string(_ab) {
		case "":
			*_cbdf = 0
		case "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c":
			*_cbdf = 1
		case "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c":
			*_cbdf = 2
		}
	}
	_gg, _cg = d.Token()
	if _cg != nil {
		return _cg
	}
	if _fcc, _agb := _gg.(_ga.EndElement); _agb && _fcc.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gg)
}

type Relationship struct{ CT_Relationship }

func (_e *CT_Relationship) UnmarshalXML(d *_ga.Decoder, start _ga.StartElement) error {
	for _, _gc := range start.Attr {
		if _gc.Name.Local == "\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065" {
			_e.TargetModeAttr.UnmarshalXMLAttr(_gc)
			continue
		}
		if _gc.Name.Local == "\u0054\u0061\u0072\u0067\u0065\u0074" {
			_ce, _ae := _gc.Value, error(nil)
			if _ae != nil {
				return _ae
			}
			_e.TargetAttr = _ce
			continue
		}
		if _gc.Name.Local == "\u0054\u0079\u0070\u0065" {
			_bb, _bf := _gc.Value, error(nil)
			if _bf != nil {
				return _bf
			}
			_e.TypeAttr = _bb
			continue
		}
		if _gc.Name.Local == "\u0049\u0064" {
			_ba, _cd := _gc.Value, error(nil)
			if _cd != nil {
				return _cd
			}
			_e.IdAttr = _ba
			continue
		}
	}
	for {
		_eb, _adb := d.Token()
		if _adb != nil {
			return _d.Errorf("p\u0061\u0072\u0073\u0069\u006e\u0067 \u0043\u0054\u005f\u0052\u0065\u006c\u0061\u0074\u0069o\u006e\u0073\u0068i\u0070:\u0020\u0025\u0073", _adb)
		}
		if _bbd, _df := _eb.(_ga.CharData); _df {
			_e.Content = string(_bbd)
		}
		if _ge, _fg := _eb.(_ga.EndElement); _fg && _ge.Name == start.Name {
			break
		}
	}
	return nil
}

func (_cbd *CT_Relationships) UnmarshalXML(d *_ga.Decoder, start _ga.StartElement) error {
_ee:
	for {
		_ed, _fc := d.Token()
		if _fc != nil {
			return _fc
		}
		switch _gbe := _ed.(type) {
		case _ga.StartElement:
			switch _gbe.Name {
			case _ga.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", Local: "\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}:
				_da := NewRelationship()
				if _bbg := d.DecodeElement(_da, &_gbe); _bbg != nil {
					return _bbg
				}
				_cbd.Relationship = append(_cbd.Relationship, _da)
			default:
				_gb.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073\u0020\u0025v", _gbe.Name)
				if _de := d.Skip(); _de != nil {
					return _de
				}
			}
		case _ga.EndElement:
			break _ee
		case _ga.CharData:
		}
	}
	return nil
}

func (_gaf ST_TargetMode) String() string {
	switch _gaf {
	case 0:
		return ""
	case 1:
		return "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c"
	case 2:
		return "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c"
	}
	return ""
}

func (_cb *CT_Relationships) MarshalXML(e *_ga.Encoder, start _ga.StartElement) error {
	e.EncodeToken(start)
	if _cb.Relationship != nil {
		_db := _ga.StartElement{Name: _ga.Name{Local: "\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}}
		for _, _fge := range _cb.Relationship {
			e.EncodeElement(_fge, _db)
		}
	}
	e.EncodeToken(_ga.EndElement{Name: start.Name})
	return nil
}

func (_ddg ST_TargetMode) MarshalXMLAttr(name _ga.Name) (_ga.Attr, error) {
	_gfe := _ga.Attr{}
	_gfe.Name = name
	switch _ddg {
	case ST_TargetModeUnset:
		_gfe.Value = ""
	case ST_TargetModeExternal:
		_gfe.Value = "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c"
	case ST_TargetModeInternal:
		_gfe.Value = "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c"
	}
	return _gfe, nil
}

// Validate validates the CT_Relationships and its children
func (_bbgb *CT_Relationships) Validate() error {
	return _bbgb.ValidateWithPath("\u0043\u0054_\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073")
}

// ValidateWithPath validates the Relationships and its children, prefixing error messages with path
func (_dd *Relationships) ValidateWithPath(path string) error {
	if _cae := _dd.CT_Relationships.ValidateWithPath(path); _cae != nil {
		return _cae
	}
	return nil
}

func (_gde *Relationships) MarshalXML(e *_ga.Encoder, start _ga.StartElement) error {
	start.Attr = append(start.Attr, _ga.Attr{Name: _ga.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s"})
	start.Attr = append(start.Attr, _ga.Attr{Name: _ga.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073"
	return _gde.CT_Relationships.MarshalXML(e, start)
}

func (_bc ST_TargetMode) Validate() error { return _bc.ValidateWithPath("") }

type Relationships struct{ CT_Relationships }

func (_gee ST_TargetMode) ValidateWithPath(path string) error {
	switch _gee {
	case 0, 1, 2:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gee))
	}
	return nil
}

const (
	ST_TargetModeUnset    ST_TargetMode = 0
	ST_TargetModeExternal ST_TargetMode = 1
	ST_TargetModeInternal ST_TargetMode = 2
)

// ValidateWithPath validates the Relationship and its children, prefixing error messages with path
func (_ceec *Relationship) ValidateWithPath(path string) error {
	if _gdd := _ceec.CT_Relationship.ValidateWithPath(path); _gdd != nil {
		return _gdd
	}
	return nil
}

func (_gbg *Relationships) UnmarshalXML(d *_ga.Decoder, start _ga.StartElement) error {
	_gbg.CT_Relationships = *NewCT_Relationships()
_dab:
	for {
		_dbe, _bfcc := d.Token()
		if _bfcc != nil {
			return _bfcc
		}
		switch _cfg := _dbe.(type) {
		case _ga.StartElement:
			switch _cfg.Name {
			case _ga.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", Local: "\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}:
				_gf := NewRelationship()
				if _fa := d.DecodeElement(_gf, &_cfg); _fa != nil {
					return _fa
				}
				_gbg.Relationship = append(_gbg.Relationship, _gf)
			default:
				_gb.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0052\u0065\u006c\u0061t\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073 \u0025\u0076", _cfg.Name)
				if _aef := d.Skip(); _aef != nil {
					return _aef
				}
			}
		case _ga.EndElement:
			break _dab
		case _ga.CharData:
		}
	}
	return nil
}

// Validate validates the Relationship and its children
func (_gce *Relationship) Validate() error {
	return _gce.ValidateWithPath("\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070")
}

// ValidateWithPath validates the CT_Relationships and its children, prefixing error messages with path
func (_adg *CT_Relationships) ValidateWithPath(path string) error {
	for _ea, _eda := range _adg.Relationship {
		if _aa := _eda.ValidateWithPath(_d.Sprintf("\u0025\u0073\u002f\u0052el\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u005b\u0025\u0064\u005d", path, _ea)); _aa != nil {
			return _aa
		}
	}
	return nil
}

func (_cad ST_TargetMode) MarshalXML(e *_ga.Encoder, start _ga.StartElement) error {
	return e.EncodeElement(_cad.String(), start)
}

func NewRelationship() *Relationship {
	_cc := &Relationship{}
	_cc.CT_Relationship = *NewCT_Relationship()
	return _cc
}

func (_bfc *Relationship) MarshalXML(e *_ga.Encoder, start _ga.StartElement) error {
	return _bfc.CT_Relationship.MarshalXML(e, start)
}

func (_fd *ST_TargetMode) UnmarshalXMLAttr(attr _ga.Attr) error {
	switch attr.Value {
	case "":
		*_fd = 0
	case "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c":
		*_fd = 1
	case "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c":
		*_fd = 2
	}
	return nil
}

// Validate validates the CT_Relationship and its children
func (_gcg *CT_Relationship) Validate() error {
	return _gcg.ValidateWithPath("\u0043T\u005fR\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070")
}

type CT_Relationships struct{ Relationship []*Relationship }

// ValidateWithPath validates the CT_Relationship and its children, prefixing error messages with path
func (_ced *CT_Relationship) ValidateWithPath(path string) error {
	if _dg := _ced.TargetModeAttr.ValidateWithPath(path + "\u002fT\u0061r\u0067\u0065\u0074\u004d\u006f\u0064\u0065\u0041\u0074\u0074\u0072"); _dg != nil {
		return _dg
	}
	return nil
}

func NewCT_Relationship() *CT_Relationship { _a := &CT_Relationship{}; return _a }

// Validate validates the Relationships and its children
func (_bd *Relationships) Validate() error {
	return _bd.ValidateWithPath("\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073")
}

func init() {
	_gb.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", "\u0043\u0054_\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073", NewCT_Relationships)
	_gb.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", "\u0043T\u005fR\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070", NewCT_Relationship)
	_gb.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", "\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073", NewRelationships)
	_gb.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", "\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070", NewRelationship)
}

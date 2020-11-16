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

package reference

import (
	_f "errors"
	_b "fmt"
	_cf "github.com/unidoc/unioffice/spreadsheet/update"
	_fb "regexp"
	_gc "strconv"
	_c "strings"
)

// IndexToColumn maps a column number to a column name (e.g. 0 = A, 1 = B, 26 = AA)
func IndexToColumn(col uint32) string {
	var _bbg [64 + 1]byte
	_dg := len(_bbg)
	_agb := col
	const _cfa = 26
	for _agb >= _cfa {
		_dg--
		_cb := _agb / _cfa
		_bbg[_dg] = byte('A' + uint(_agb-_cb*_cfa))
		_agb = _cb - 1
	}
	_dg--
	_bbg[_dg] = byte('A' + uint(_agb))
	return string(_bbg[_dg:])
}

// ParseRangeReference splits a range reference of the form "A1:B5" into its
// components.
func ParseRangeReference(s string) (_eb, _da CellReference, _ab error) {
	_ggc, _de, _ab := _gg(s)
	if _ab != nil {
		return CellReference{}, CellReference{}, _ab
	}
	_cfb := _c.Split(_de, "\u003a")
	if len(_cfb) != 2 {
		return CellReference{}, CellReference{}, _f.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006eg\u0065\u0020\u0066\u006frm\u0061\u0074")
	}
	if _ggc != "" {
		_cfb[0] = _ggc + "\u0021" + _cfb[0]
		_cfb[1] = _ggc + "\u0021" + _cfb[1]
	}
	_eg, _ab := ParseCellReference(_cfb[0])
	if _ab != nil {
		return CellReference{}, CellReference{}, _ab
	}
	_dd, _ab := ParseCellReference(_cfb[1])
	if _ab != nil {
		return CellReference{}, CellReference{}, _ab
	}
	return _eg, _dd, nil
}

// ParseColumnRangeReference splits a range reference of the form "A:B" into its
// components.
func ParseColumnRangeReference(s string) (_cbe, _ed ColumnReference, _bfa error) {
	_ac := ""
	_aa := _c.Split(s, "\u0021")
	if len(_aa) == 2 {
		_ac = _aa[0]
		s = _aa[1]
	}
	_ca := _c.Split(s, "\u003a")
	if len(_ca) != 2 {
		return ColumnReference{}, ColumnReference{}, _f.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006eg\u0065\u0020\u0066\u006frm\u0061\u0074")
	}
	if _ac != "" {
		_ca[0] = _ac + "\u0021" + _ca[0]
		_ca[1] = _ac + "\u0021" + _ca[1]
	}
	_agef, _bfa := ParseColumnReference(_ca[0])
	if _bfa != nil {
		return ColumnReference{}, ColumnReference{}, _bfa
	}
	_abf, _bfa := ParseColumnReference(_ca[1])
	if _bfa != nil {
		return ColumnReference{}, ColumnReference{}, _bfa
	}
	return _agef, _abf, nil
}

var _bga = _fb.MustCompile("^\u005b\u0061\u002d\u007aA-\u005a]\u0028\u005b\u0061\u002d\u007aA\u002d\u005a\u005d\u003f\u0029\u0024")

// CellReference is a parsed reference to a cell.  Input is of the form 'A1',
// '$C$2', etc.
type CellReference struct {
	RowIdx         uint32
	ColumnIdx      uint32
	Column         string
	AbsoluteColumn bool
	AbsoluteRow    bool
	SheetName      string
}

// ColumnReference is a parsed reference to a column.  Input is of the form 'A',
// '$C', etc.
type ColumnReference struct {
	ColumnIdx      uint32
	Column         string
	AbsoluteColumn bool
	SheetName      string
}

// Update updates reference to point one of the neighboring cells with respect to the update type after removing a row/column.
func (_ge *CellReference) Update(updateType _cf.UpdateAction) *CellReference {
	switch updateType {
	case _cf.UpdateActionRemoveColumn:
		_a := _ge
		_a.ColumnIdx = _ge.ColumnIdx - 1
		_a.Column = IndexToColumn(_a.ColumnIdx)
		return _a
	default:
		return _ge
	}
}

// Update updates reference to point one of the neighboring columns with respect to the update type after removing a row/column.
func (_ag *ColumnReference) Update(updateType _cf.UpdateAction) *ColumnReference {
	switch updateType {
	case _cf.UpdateActionRemoveColumn:
		_ec := _ag
		_ec.ColumnIdx = _ag.ColumnIdx - 1
		_ec.Column = IndexToColumn(_ec.ColumnIdx)
		return _ec
	default:
		return _ag
	}
}

func _gg(_gcg string) (string, string, error) {
	_dbe := ""
	_eef := _c.LastIndex(_gcg, "\u0021")
	if _eef > -1 {
		_dbe = _gcg[:_eef]
		_gcg = _gcg[_eef+1:]
		if _dbe == "" {
			return "", "", _f.New("\u0049n\u0076a\u006c\u0069\u0064\u0020\u0073h\u0065\u0065t\u0020\u006e\u0061\u006d\u0065")
		}
	}
	return _dbe, _gcg, nil
}

// String returns a string representation of CellReference.
func (_gb CellReference) String() string {
	_e := make([]byte, 0, 4)
	if _gb.AbsoluteColumn {
		_e = append(_e, '$')
	}
	_e = append(_e, _gb.Column...)
	if _gb.AbsoluteRow {
		_e = append(_e, '$')
	}
	_e = _gc.AppendInt(_e, int64(_gb.RowIdx), 10)
	return string(_e)
}

// ColumnToIndex maps a column to a zero based index (e.g. A = 0, B = 1, AA = 26)
func ColumnToIndex(col string) uint32 {
	col = _c.ToUpper(col)
	_age := uint32(0)
	for _, _ccf := range col {
		_age *= 26
		_age += uint32(_ccf - 'A' + 1)
	}
	return _age - 1
}

// ParseColumnReference parses a column reference of the form 'Sheet1!A' and splits it
// into sheet name and column segments.
func ParseColumnReference(s string) (ColumnReference, error) {
	s = _c.TrimSpace(s)
	if len(s) < 1 {
		return ColumnReference{}, _f.New("\u0063\u006f\u006c\u0075\u006d\u006e \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065\u0020a\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0063\u0068a\u0072a\u0063\u0074\u0065\u0072")
	}
	_cc := ColumnReference{}
	_af, _bb, _ae := _gg(s)
	if _ae != nil {
		return ColumnReference{}, _ae
	}
	if _af != "" {
		_cc.SheetName = _af
	}
	if _bb[0] == '$' {
		_cc.AbsoluteColumn = true
		_bb = _bb[1:]
	}
	if !_bga.MatchString(_bb) {
		return ColumnReference{}, _f.New("\u0063\u006f\u006c\u0075\u006dn\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0041\u0020\u0061\u006e\u0064\u0020\u005a\u005a")
	}
	_cc.Column = _bb
	_cc.ColumnIdx = ColumnToIndex(_cc.Column)
	return _cc, nil
}

// ParseCellReference parses a cell reference of the form 'A10' and splits it
// into column/row segments.
func ParseCellReference(s string) (CellReference, error) {
	s = _c.TrimSpace(s)
	if len(s) < 2 {
		return CellReference{}, _f.New("\u0063\u0065\u006c\u006c\u0020\u0072\u0065\u0066e\u0072\u0065\u006ece\u0020\u006d\u0075\u0073\u0074\u0020h\u0061\u0076\u0065\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0074\u0077o\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074e\u0072\u0073")
	}
	_bf := CellReference{}
	_ba, _d, _bg := _gg(s)
	if _bg != nil {
		return CellReference{}, _bg
	}
	if _ba != "" {
		_bf.SheetName = _ba
	}
	if s[0] == '$' {
		_bf.AbsoluteColumn = true
		_d = _d[1:]
	}
	_ee := -1
_db:
	for _fe := 0; _fe < len(_d); _fe++ {
		switch {
		case _d[_fe] >= '0' && _d[_fe] <= '9' || _d[_fe] == '$':
			_ee = _fe
			break _db
		}
	}
	switch _ee {
	case 0:
		return CellReference{}, _b.Errorf("\u006e\u006f\u0020\u006cet\u0074\u0065\u0072\u0020\u0070\u0072\u0065\u0066\u0069\u0078\u0020\u0069\u006e\u0020%\u0073", _d)
	case -1:
		return CellReference{}, _b.Errorf("\u006eo\u0020d\u0069\u0067\u0069\u0074\u0073\u0020\u0069\u006e\u0020\u0025\u0073", _d)
	}
	_bf.Column = _d[0:_ee]
	if _d[_ee] == '$' {
		_bf.AbsoluteRow = true
		_ee++
	}
	_bf.ColumnIdx = ColumnToIndex(_bf.Column)
	_fd, _bg := _gc.ParseUint(_d[_ee:], 10, 32)
	if _bg != nil {
		return CellReference{}, _b.Errorf("e\u0072\u0072\u006f\u0072 p\u0061r\u0073\u0069\u006e\u0067\u0020r\u006f\u0077\u003a\u0020\u0025\u0073", _bg)
	}
	_bf.RowIdx = uint32(_fd)
	return _bf, nil
}

// String returns a string representation of ColumnReference.
func (_gbe ColumnReference) String() string {
	_bc := make([]byte, 0, 4)
	if _gbe.AbsoluteColumn {
		_bc = append(_bc, '$')
	}
	_bc = append(_bc, _gbe.Column...)
	return string(_bc)
}

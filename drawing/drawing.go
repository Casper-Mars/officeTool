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

package drawing

import (
	_c "github.com/Casper-Mars/officeTool"
	_g "github.com/Casper-Mars/officeTool/color"
	_d "github.com/Casper-Mars/officeTool/measurement"
	_a "github.com/Casper-Mars/officeTool/schema/soo/dml"
)

// X returns the inner wrapped XML type.
func (_ca Paragraph) X() *_a.CT_TextParagraph { return _ca._fe }

// SetBold controls the bolding of a run.
func (_egg RunProperties) SetBold(b bool) { _egg._bbb.BAttr = _c.Bool(b) }

// ParagraphProperties allows controlling paragraph properties.
type ParagraphProperties struct {
	_eb *_a.CT_TextParagraphProperties
}

// SetHeight sets the height of the shape.
func (_aef ShapeProperties) SetHeight(h _d.Distance) {
	_aef.ensureXfrm()
	if _aef._fgg.Xfrm.Ext == nil {
		_aef._fgg.Xfrm.Ext = _a.NewCT_PositiveSize2D()
	}
	_aef._fgg.Xfrm.Ext.CyAttr = int64(h / _d.EMU)
}

func (_caf ShapeProperties) LineProperties() LineProperties {
	if _caf._fgg.Ln == nil {
		_caf._fgg.Ln = _a.NewCT_LineProperties()
	}
	return LineProperties{_caf._fgg.Ln}
}

// X returns the inner wrapped XML type.
func (_gbf ShapeProperties) X() *_a.CT_ShapeProperties { return _gbf._fgg }

// SetPosition sets the position of the shape.
func (_bba ShapeProperties) SetPosition(x, y _d.Distance) {
	_bba.ensureXfrm()
	if _bba._fgg.Xfrm.Off == nil {
		_bba._fgg.Xfrm.Off = _a.NewCT_Point2D()
	}
	_bba._fgg.Xfrm.Off.XAttr.ST_CoordinateUnqualified = _c.Int64(int64(x / _d.EMU))
	_bba._fgg.Xfrm.Off.YAttr.ST_CoordinateUnqualified = _c.Int64(int64(y / _d.EMU))
}

// SetLevel sets the level of indentation of a paragraph.
func (_dd ParagraphProperties) SetLevel(idx int32) { _dd._eb.LvlAttr = _c.Int32(idx) }

// SetWidth sets the width of the shape.
func (_dgb ShapeProperties) SetWidth(w _d.Distance) {
	_dgb.ensureXfrm()
	if _dgb._fgg.Xfrm.Ext == nil {
		_dgb._fgg.Xfrm.Ext = _a.NewCT_PositiveSize2D()
	}
	_dgb._fgg.Xfrm.Ext.CxAttr = int64(w / _d.EMU)
}

// AddRun adds a new run to a paragraph.
func (_e Paragraph) AddRun() Run {
	_gf := MakeRun(_a.NewEG_TextRun())
	_e._fe.EG_TextRun = append(_e._fe.EG_TextRun, _gf.X())
	return _gf
}

// RunProperties controls the run properties.
type RunProperties struct {
	_bbb *_a.CT_TextCharacterProperties
}

// SetJoin sets the line join style.
func (_be LineProperties) SetJoin(e LineJoin) {
	_be._dc.Round = nil
	_be._dc.Miter = nil
	_be._dc.Bevel = nil
	switch e {
	case LineJoinRound:
		_be._dc.Round = _a.NewCT_LineJoinRound()
	case LineJoinBevel:
		_be._dc.Bevel = _a.NewCT_LineJoinBevel()
	case LineJoinMiter:
		_be._dc.Miter = _a.NewCT_LineJoinMiterProperties()
	}
}

// SetBulletChar sets the bullet character for the paragraph.
func (_gga ParagraphProperties) SetBulletChar(c string) {
	if c == "" {
		_gga._eb.BuChar = nil
	} else {
		_gga._eb.BuChar = _a.NewCT_TextCharBullet()
		_gga._eb.BuChar.CharAttr = c
	}
}

// SetFlipHorizontal controls if the shape is flipped horizontally.
func (_ge ShapeProperties) SetFlipHorizontal(b bool) {
	_ge.ensureXfrm()
	if !b {
		_ge._fgg.Xfrm.FlipHAttr = nil
	} else {
		_ge._fgg.Xfrm.FlipHAttr = _c.Bool(true)
	}
}

// LineJoin is the type of line join
type LineJoin byte

// Properties returns the paragraph properties.
func (_ee Paragraph) Properties() ParagraphProperties {
	if _ee._fe.PPr == nil {
		_ee._fe.PPr = _a.NewCT_TextParagraphProperties()
	}
	return MakeParagraphProperties(_ee._fe.PPr)
}

func (_ggd ShapeProperties) ensureXfrm() {
	if _ggd._fgg.Xfrm == nil {
		_ggd._fgg.Xfrm = _a.NewCT_Transform2D()
	}
}

// SetAlign controls the paragraph alignment
func (_df ParagraphProperties) SetAlign(a _a.ST_TextAlignType) { _df._eb.AlgnAttr = a }

// SetFlipVertical controls if the shape is flipped vertically.
func (_aa ShapeProperties) SetFlipVertical(b bool) {
	_aa.ensureXfrm()
	if !b {
		_aa._fgg.Xfrm.FlipVAttr = nil
	} else {
		_aa._fgg.Xfrm.FlipVAttr = _c.Bool(true)
	}
}

// SetWidth sets the line width, MS products treat zero as the minimum width
// that can be displayed.
func (_ab LineProperties) SetWidth(w _d.Distance) { _ab._dc.WAttr = _c.Int32(int32(w / _d.EMU)) }

// SetFont controls the font of a run.
func (_bf RunProperties) SetFont(s string) {
	_bf._bbb.Latin = _a.NewCT_TextFont()
	_bf._bbb.Latin.TypefaceAttr = s
}

// MakeRun constructs a new Run wrapper.
func MakeRun(x *_a.EG_TextRun) Run { return Run{x} }

// X returns the inner wrapped XML type.
func (_gd ParagraphProperties) X() *_a.CT_TextParagraphProperties { return _gd._eb }

type ShapeProperties struct{ _fgg *_a.CT_ShapeProperties }

// SetSize sets the width and height of the shape.
func (_eba ShapeProperties) SetSize(w, h _d.Distance) { _eba.SetWidth(w); _eba.SetHeight(h) }

// MakeRunProperties constructs a new RunProperties wrapper.
func MakeRunProperties(x *_a.CT_TextCharacterProperties) RunProperties { return RunProperties{x} }

func (_bg ShapeProperties) clearFill() {
	_bg._fgg.NoFill = nil
	_bg._fgg.BlipFill = nil
	_bg._fgg.GradFill = nil
	_bg._fgg.GrpFill = nil
	_bg._fgg.SolidFill = nil
	_bg._fgg.PattFill = nil
}

func (_bb LineProperties) clearFill() {
	_bb._dc.NoFill = nil
	_bb._dc.GradFill = nil
	_bb._dc.SolidFill = nil
	_bb._dc.PattFill = nil
}

func MakeShapeProperties(x *_a.CT_ShapeProperties) ShapeProperties { return ShapeProperties{x} }

// X returns the inner wrapped XML type.
func (_eg Run) X() *_a.EG_TextRun { return _eg._gb }

const (
	LineJoinRound LineJoin = iota
	LineJoinBevel
	LineJoinMiter
)

// SetSize sets the font size of the run text
func (_ag RunProperties) SetSize(sz _d.Distance) {
	_ag._bbb.SzAttr = _c.Int32(int32(sz / _d.HundredthPoint))
}

// Run is a run within a paragraph.
type Run struct{ _gb *_a.EG_TextRun }

func (_dg ShapeProperties) SetNoFill() {
	_dg.clearFill()
	_dg._fgg.NoFill = _a.NewCT_NoFillProperties()
}

// SetText sets the run's text contents.
func (_da Run) SetText(s string) {
	_da._gb.Br = nil
	_da._gb.Fld = nil
	if _da._gb.R == nil {
		_da._gb.R = _a.NewCT_RegularTextRun()
	}
	_da._gb.R.T = s
}

func (_bc ShapeProperties) SetSolidFill(c _g.Color) {
	_bc.clearFill()
	_bc._fgg.SolidFill = _a.NewCT_SolidColorFillProperties()
	_bc._fgg.SolidFill.SrgbClr = _a.NewCT_SRgbColor()
	_bc._fgg.SolidFill.SrgbClr.ValAttr = *c.AsRGBString()
}

// AddBreak adds a new line break to a paragraph.
func (_gg Paragraph) AddBreak() {
	_gfe := _a.NewEG_TextRun()
	_gfe.Br = _a.NewCT_TextLineBreak()
	_gg._fe.EG_TextRun = append(_gg._fe.EG_TextRun, _gfe)
}

func (_f LineProperties) SetSolidFill(c _g.Color) {
	_f.clearFill()
	_f._dc.SolidFill = _a.NewCT_SolidColorFillProperties()
	_f._dc.SolidFill.SrgbClr = _a.NewCT_SRgbColor()
	_f._dc.SolidFill.SrgbClr.ValAttr = *c.AsRGBString()
}

type LineProperties struct{ _dc *_a.CT_LineProperties }

// SetGeometry sets the shape type of the shape
func (_fb ShapeProperties) SetGeometry(g _a.ST_ShapeType) {
	if _fb._fgg.PrstGeom == nil {
		_fb._fgg.PrstGeom = _a.NewCT_PresetGeometry2D()
	}
	_fb._fgg.PrstGeom.PrstAttr = g
}

// SetBulletFont controls the font for the bullet character.
func (_ae ParagraphProperties) SetBulletFont(f string) {
	if f == "" {
		_ae._eb.BuFont = nil
	} else {
		_ae._eb.BuFont = _a.NewCT_TextFont()
		_ae._eb.BuFont.TypefaceAttr = f
	}
}

// Paragraph is a paragraph within a document.
type Paragraph struct{ _fe *_a.CT_TextParagraph }

// X returns the inner wrapped XML type.
func (_cc LineProperties) X() *_a.CT_LineProperties { return _cc._dc }

func (_abe LineProperties) SetNoFill() {
	_abe.clearFill()
	_abe._dc.NoFill = _a.NewCT_NoFillProperties()
}

// MakeParagraphProperties constructs a new ParagraphProperties wrapper.
func MakeParagraphProperties(x *_a.CT_TextParagraphProperties) ParagraphProperties {
	return ParagraphProperties{x}
}

// MakeParagraph constructs a new paragraph wrapper.
func MakeParagraph(x *_a.CT_TextParagraph) Paragraph { return Paragraph{x} }

// SetNumbered controls if bullets are numbered or not.
func (_fc ParagraphProperties) SetNumbered(scheme _a.ST_TextAutonumberScheme) {
	if scheme == _a.ST_TextAutonumberSchemeUnset {
		_fc._eb.BuAutoNum = nil
	} else {
		_fc._eb.BuAutoNum = _a.NewCT_TextAutonumberBullet()
		_fc._eb.BuAutoNum.TypeAttr = scheme
	}
}

// SetSolidFill controls the text color of a run.
func (_fg RunProperties) SetSolidFill(c _g.Color) {
	_fg._bbb.NoFill = nil
	_fg._bbb.BlipFill = nil
	_fg._bbb.GradFill = nil
	_fg._bbb.GrpFill = nil
	_fg._bbb.PattFill = nil
	_fg._bbb.SolidFill = _a.NewCT_SolidColorFillProperties()
	_fg._bbb.SolidFill.SrgbClr = _a.NewCT_SRgbColor()
	_fg._bbb.SolidFill.SrgbClr.ValAttr = *c.AsRGBString()
}

// Properties returns the run's properties.
func (_cd Run) Properties() RunProperties {
	if _cd._gb.R == nil {
		_cd._gb.R = _a.NewCT_RegularTextRun()
	}
	if _cd._gb.R.RPr == nil {
		_cd._gb.R.RPr = _a.NewCT_TextCharacterProperties()
	}
	return RunProperties{_cd._gb.R.RPr}
}

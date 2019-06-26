package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/vml/internal"
)

//FillType is direct mapping of ST_FillType
type FillType string //enum

//FillMethod is direct mapping of ST_FillMethod
type FillMethod string //enum

//Fill is direct mapping of CT_Fill
type Fill struct {
	XMLName          xml.Name    `xml:"fill" namespace:"v"`
	AlignShape       *bool       `xml:"alignshape,attr,omitempty"`
	AltHRef          string      `xml:"althref,attr,omitempty" namespace:"o"`
	Angle            *int        `xml:"angle,attr,omitempty"`
	Aspect           ImageAspect `xml:"aspect,attr,omitempty"`
	Color            string      `xml:"color,attr,omitempty"`
	Color2           string      `xml:"color2,attr,omitempty"`
	Colors           string      `xml:"colors,attr,omitempty"`
	DetectMouseClick *bool       `xml:"detectmouseclick,attr,omitempty" namespace:"o"`
	Focus            string      `xml:"focus,attr,omitempty"`
	FocusPosition    string      `xml:"focusposition,attr,omitempty"`
	FocusSize        string      `xml:"focussize,attr,omitempty"`
	HRef             string      `xml:"href,attr,omitempty" namespace:"o"`
	ID               string      `xml:"id,attr,omitempty"`
	Method           FillMethod  `xml:"method,attr,omitempty"`
	On               *bool       `xml:"on,attr,omitempty"`
	Opacity          string      `xml:"opacity,attr,omitempty"`
	Opacity2         string      `xml:"opacity2,attr,omitempty" namespace:"o"`
	Origin           string      `xml:"origin,attr,omitempty"`
	Position         string      `xml:"position,attr,omitempty"`
	Recolor          *bool       `xml:"recolor,attr,omitempty"`
	Rotate           *bool       `xml:"rotate,attr,omitempty"`
	RelID            string      `xml:"relid,attr,omitempty" namespace:"o"`
	Size             string      `xml:"size,attr,omitempty"`
	Src              string      `xml:"src,attr,omitempty"`
	Title            string      `xml:"title,attr,omitempty" namespace:"o"`
	Type             FillType    `xml:"type,attr,omitempty"`
	Extended         *FillExtended

	//FIXME: 'r:id' conflicts with 'id'
	//RID              string            `xml:"r id,attr,omitempty"`
}

type FillExtended struct {
	XMLName xml.Name `xml:"fill" namespace:"o"`
	Type    FillType `xml:"type,attr,omitempty"`
	Ext     Ext      `xml:"ext,attr,omitempty" namespace:"v"`
}

func (s *Fill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.Encode(s, e)
}

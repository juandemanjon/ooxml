// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import "github.com/plandem/ooxml/ml"

//BlipFillProperties is a direct mapping of XSD CT_BlipFillProperties
type BlipFillProperties struct {
	Blip            *Blip `xml:"blip,omitempty"`
	Dpi             int   `xml:"dpi,attr,omitempty"`
	RotateWithShape bool  `xml:"rotWithShape,attr,omitempty"`
	ml.ReservedElements
}

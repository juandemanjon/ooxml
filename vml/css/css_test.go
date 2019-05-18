package css_test

import (
	"encoding/xml"
	"github.com/plandem/ooxml/vml/css"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCSS(t *testing.T) {
	data := "position:absolute;margin-left:59.25pt;margin-top:1.50cm;width:96px;height:55px;z-index:1;visibility:hidden"
	decoded := css.Decode(data)

	require.Equal(t, css.Style{
		Position:   css.PositionAbsolute,
		MarginLeft: css.NumberPt(59.25),
		MarginTop:  css.NumberCm(1.5),
		Width:      css.NumberPx(96),
		Height:     css.NumberPx(55),
		ZIndex:     1,
		Visible:    css.VisibilityHidden,
	}, decoded)

	encoded := decoded.Encode()
	require.Equal(t, data, encoded)
	require.Equal(t, decoded, css.Decode(encoded))
}

func TestXml(t *testing.T) {
	type Entity struct {
		Style  css.Style  `xml:"style,attr,omitempty"`
		PStyle *css.Style `xml:"p_style,attr,omitempty"`
	}

	//empty
	entity := Entity{Style: css.Style{}, PStyle: &css.Style{}}
	encoded, err := xml.Marshal(&entity)
	require.Empty(t, err)
	require.Equal(t, `<Entity></Entity>`, string(encoded))

	//encode
	s := css.Style{
		Position:   css.PositionAbsolute,
		MarginLeft: css.NumberPt(59.25),
		MarginTop:  css.NumberPt(1.5),
		Width:      css.NumberPx(96),
		Height:     css.NumberPx(55),
		ZIndex:     1,
		Visible:    css.VisibilityHidden,
	}
	entity = Entity{Style: s, PStyle: &s}
	encoded, err = xml.Marshal(&entity)

	require.Empty(t, err)
	require.Equal(t, `<Entity style="position:absolute;margin-left:59.25pt;margin-top:1.50pt;width:96px;height:55px;z-index:1;visibility:hidden" p_style="position:absolute;margin-left:59.25pt;margin-top:1.50pt;width:96px;height:55px;z-index:1;visibility:hidden"></Entity>`, string(encoded))

	//decode
	var decoded Entity
	err = xml.Unmarshal(encoded, &decoded)
	require.Empty(t, err)

	require.Equal(t, entity, decoded)
}
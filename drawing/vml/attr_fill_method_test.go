package vml_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml/drawing/vml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFillMethod(t *testing.T) {
	type Entity struct {
		Attribute vml.FillMethod `xml:"attribute,attr,omitempty"`
	}

	list := map[vml.FillMethod]string{
		vml.FillMethod(0):         "",
		vml.FillMethodNone:        vml.FillMethodNone.String(),
		vml.FillMethodAny:         vml.FillMethodAny.String(),
		vml.FillMethodLinear:      vml.FillMethodLinear.String(),
		vml.FillMethodLinearSigma: vml.FillMethodLinearSigma.String(),
		vml.FillMethodSigma:       vml.FillMethodSigma.String(),
	}

	for k, v := range list {
		t.Run(v, func(tt *testing.T) {
			entity := Entity{Attribute: k}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if k == 0 {
				require.Equal(tt, `<Entity></Entity>`, string(encoded))
			} else {
				require.Equal(tt, fmt.Sprintf(`<Entity attribute="%s"></Entity>`, v), string(encoded))
			}

			var decoded Entity
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
			require.Equal(tt, v, decoded.Attribute.String())
		})
	}
}
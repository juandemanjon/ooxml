// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ml_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProperty(t *testing.T) {
	type Element struct {
		Property  ml.Property  `xml:"property"`
		PProperty *ml.Property `xml:"p_property"`
	}

	list := map[string]ml.Property{
		"1":      "1",
		"aaa":    "aaa",
		" test ": " test ",
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Element{Property: v, PProperty: &v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			require.Equal(tt, fmt.Sprintf(`<Element><property val="%s"></property><p_property val="%s"></p_property></Element>`, s, s), string(encoded))

			var decoded Element
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
		})
	}
}

func TestPropertyInt(t *testing.T) {
	type Element struct {
		Property  ml.PropertyInt  `xml:"property"`
		PProperty *ml.PropertyInt `xml:"p_property"`
	}

	list := map[string]ml.PropertyInt{
		"-2": -2,
		"-1": -1,
		"0":  0,
		"1":  1,
		"2":  2,
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Element{Property: v, PProperty: &v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			require.Equal(tt, fmt.Sprintf(`<Element><property val="%s"></property><p_property val="%s"></p_property></Element>`, s, s), string(encoded))

			var decoded Element
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
		})
	}
}

func TestPropertyDouble(t *testing.T) {
	type Element struct {
		Property  ml.PropertyDouble  `xml:"property"`
		PProperty *ml.PropertyDouble `xml:"p_property"`
	}

	list := map[string]ml.PropertyDouble{
		"-2.2": -2.2,
		"-1":   -1.0,
		"0":    0.0,
		"1":    1.0,
		"2.2":  2.2,
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Element{Property: v, PProperty: &v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			require.Equal(tt, fmt.Sprintf(`<Element><property val="%s"></property><p_property val="%s"></p_property></Element>`, s, s), string(encoded))

			var decoded Element
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
		})
	}
}

func TestPropertyBool(t *testing.T) {
	type Element struct {
		Property  ml.PropertyBool  `xml:"property"`
		PProperty *ml.PropertyBool `xml:"p_property"`
	}

	list := map[string]ml.PropertyBool{
		"true":  true,
		"false": false,
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Element{Property: v, PProperty: &v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			require.Equal(tt, fmt.Sprintf(`<Element><property val="%s"></property><p_property val="%s"></p_property></Element>`, s, s), string(encoded))

			var decoded Element
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
		})
	}

	var decoded Element
	err := xml.Unmarshal([]byte(`<Element><property/></Element>`), &decoded)
	require.Empty(t, err)
	require.Equal(t, Element{Property: true}, decoded)
}

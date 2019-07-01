package ml

import "encoding/xml"

//Reserved is special type that catches all inner content AS IS to save original information - used to mark 'non implemented' elements
type Reserved struct {
	XMLName  xml.Name
	InnerXML string `xml:",innerxml"`
	ReservedAttributes
}

//ReservedAttributes is a special type that catches all not captured attributes AS IS to save original information - used to mark 'non implemented' attributes
type ReservedAttributes struct {
	Attrs []xml.Attr `xml:",any,attr"`
}

//ReservedElements is a special type that catches all not captured nested elements AS IS to save original information - used to mark 'non implemented' elements
type ReservedElements struct {
	Nodes []Reserved `xml:",any"`
}

//ResolveNamespacePrefixes transforms namespaces into namespaces prefixes
func (r ReservedAttributes) ResolveNamespacePrefixes() {
	for i, attr := range r.Attrs {
		r.Attrs[i].Name = ApplyNamespacePrefix(attr.Name.Space, attr.Name)
	}
}

//ResolveNamespacePrefixes tries to resolve namespace and apply prefix for it for all reserved elements
func (r ReservedElements) ResolveNamespacePrefixes() {
	for i, node := range r.Nodes {
		r.Nodes[i].XMLName = ApplyNamespacePrefix(node.XMLName.Space, node.XMLName)
		node.ResolveNamespacePrefixes()
	}
}

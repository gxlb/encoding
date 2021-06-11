package xml

import (
	"fmt"
	"strconv"

	"github.com/beevik/etree"
)

// Decoder implements xml decoder
type Decoder struct{}

// Decode unmarshal xml to universal Go data
func (dec *Decoder) Decode(data []byte) (interface{}, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(data); err != nil {
		return nil, err
	}
	return dec.visit(doc.Root())
}

// getAttr return Attr of an xml element with given name
func getAttr(name string, v *etree.Element) string {
	for i := 0; i < len(v.Attr); i++ {
		if p := &v.Attr[i]; p.Key == name {
			return p.Value
		}
	}
	return ""
}

// visit access an xml element recursively
func (dec *Decoder) visit(v *etree.Element) (interface{}, error) {
	switch a := getAttr(valueType, v); a {
	case ValueBool:
		return strconv.ParseBool(v.Text())
	case ValueNumber:
		return strconv.ParseFloat(v.Text(), 64)
	case ValueString:
		return v.Text(), nil
	case ValueArray:
		c := v.ChildElements()
		var s = make([]interface{}, 0, len(c))
		for i := 0; i < len(c); i++ {
			vv := c[i]
			d, err := dec.visit(vv)
			if err != nil {
				return nil, err
			}
			s = append(s, d)
		}
		return s, nil
	case ValueMap:
		c := v.ChildElements()
		var s = make(map[string]interface{})
		for i := 0; i < len(c); i++ {
			vv := c[i]
			d, err := dec.visit(vv)
			if err != nil {
				return nil, err
			}
			s[v.Tag] = d
		}
		return s, nil
	case "":
		c := v.ChildElements()
		if len(c) == 0 {
			t := v.Text()
			// if d, err := strconv.ParseFloat(t, 64); err == nil {
			// 	return d, nil
			// }
			// if d, err := strconv.ParseBool(t); err == nil {
			// 	return d, nil
			// }
			if t == "null" {
				return nil, nil
			}
			return t, nil
		} else {
			var s = make(map[string]interface{})
			for i := 0; i < len(c); i++ {
				vv := c[i]
				d, err := dec.visit(vv)
				if err != nil {
					return nil, err
				}
				s[vv.Tag] = d
			}
			return s, nil
		}
	default:
	}

	return nil, fmt.Errorf("unsupported node %#v", v)
}

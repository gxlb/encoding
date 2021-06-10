package xml

import (
	"fmt"

	"github.com/beevik/etree"
)

type Decoder struct{}

func (dec *Decoder) Decode(d interface{}, data []byte) error {
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(data); err != nil {
		return err
	}
	return dec.visit(doc.Root())
}

func (dec *Decoder) visit(v *etree.Element) error {
	fmt.Println(v.Tag, v.Text(), v.Attr)

	c := v.ChildElements()
	for i := 0; i < len(c); i++ {
		vv := c[i]
		if err := dec.visit(vv); err != nil {
			return err
		}
	}
	return nil
}

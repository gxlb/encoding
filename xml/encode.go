package xml

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/beevik/etree"
)

type Encoder struct{}

func (e *Encoder) Encode(v interface{}, buf []byte) ([]byte, error) {
	d := etree.NewDocument()
	d.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	root := d.CreateElement("root")
	e.encodeValue(root, "data", reflect.ValueOf(v))

	b := bytes.NewBuffer(buf)
	_, err := d.WriteTo(b)
	return b.Bytes(), err
}

func (e *Encoder) value(parent *etree.Element, name string, v reflect.Value) error {
	switch k := v.Kind(); k {
	case reflect.Int, reflect.Uint, reflect.Float32, reflect.Float64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.String, reflect.Bool:
		elem := parent.CreateElement(name)
		elem.SetText(fmt.Sprintf("%v", v.Interface()))
		return nil

	case reflect.Array, reflect.Slice:
		l := v.Len()
		elem := parent.CreateElement(name)
		for i := 0; i < l; i++ {
			if err := e.value(elem, "elem", v.Index(i)); err != nil {
				return err
			}
		}
		return nil

	case reflect.Struct:
		t := v.Type()
		elem := parent.CreateElement(name)
		for i, n := 0, t.NumField(); i < n; i++ {
			f := t.Field(i)
			fv := v.Field(i)
			if err := e.value(elem, f.Name, fv); err != nil {
				return err
			}
		}
		return nil

	case reflect.Map:
		t := v.Type()
		kt := t.Key()
		vt := t.Elem()
		if kt.Kind() != reflect.String {
			return fmt.Errorf("unsupported map key %s", kt.String())
		}
		keys := v.MapKeys()
		elem := parent.CreateElement(name)
		for i, key := range keys {
			if err := e.value(elem, key, v.MapIndex(key)); err != nil {
				return err
			}
		}
		return nil

	case reflect.Ptr:
		if !v.IsNil() {
			return e.value(parent, name, v.Elem())
		}
		return nil

	default:
		// do nothing
	}
	return fmt.Errorf("unsupported type %s", v.Type().String())
}
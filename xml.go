package encoding

import (
	"polyapi/internal/encoding/xml"
)

//-------------------------

// ToXml encode a Go data to xml string.
func ToXml(v interface{}, pretty bool) (string, error) {
	enc := xml.Encoder{}
	b, err := enc.Encode(v, nil, pretty)
	return unsafeByteString(b), err
}

// MustToXml encode a Go data to xml string.
// It panic if error.
func MustToXml(v interface{}, pretty bool) string {
	s, err := ToXml(v, pretty)
	if err != nil {
		panic(err)
	}
	return s
}

// FromXml decode a xml string to universal Go data.
func FromXml(s string) (interface{}, error) {
	dec := xml.Decoder{}
	d, err := dec.Decode(unsafeStringBytes(s))
	return d, err
}

// MustFromXml decode a xml string to universal Go data.
// It panic if error.
func MustFromXml(s string) interface{} {
	d, err := FromXml(s)
	if err != nil {
		panic(err)
	}
	return d
}

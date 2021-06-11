package encoding

import (
	"encoding/json"
	"fmt"
)

//-------------------------

// ToJson encode a Go data to json string
func ToJson(v interface{}, pretty bool) (string, error) {
	var b []byte
	var err error
	if pretty {
		b, err = json.MarshalIndent(v, "", "  ")
	} else {
		b, err = json.Marshal(v)
	}

	return unsafeByteString(b), err
}

// MustToJson encode a Go data to json string.
// It panic if error.
func MustToJson(v interface{}, pretty bool) string {
	s, err := ToJson(v, pretty)
	if err != nil {
		panic(err)
	}
	return s
}

// FromJson decode a json string to universal Go data
func FromJson(s string) (interface{}, error) {
	var d struct {
		D interface{}
	}
	ss := fmt.Sprintf(`{"D":%s}`, s)
	err := json.Unmarshal(unsafeStringBytes(ss), &d)
	return d.D, err
}

// MustFromJson decode a json string to universal Go data.
// It panic if error.
func MustFromJson(s string) interface{} {
	d, err := FromJson(s)
	if err != nil {
		panic(err)
	}
	return d
}

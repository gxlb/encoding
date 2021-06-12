package encoding

import (
	yaml "gopkg.in/yaml.v2"
)

//-------------------------

// ToYaml encode a Go data to yaml string.
func ToYaml(v interface{}, pretty bool) (string, error) {
	b, err := yaml.Marshal(v)
	return unsafeByteString(b), err
}

// MustToYaml encode a Go data to yaml string.
// It panic if error.
func MustToYaml(v interface{}, pretty bool) string {
	s, err := ToYaml(v, pretty)
	if err != nil {
		panic(err)
	}
	return s
}

// FromYaml decode a yaml string to universal Go data.
func FromYaml(s string) (interface{}, error) {
	var d interface{}
	err := yaml.Unmarshal(unsafeStringBytes(s), &d) // decode to *interface{}
	return d, err
}

// MustFromYaml decode a yaml string to universal Go data.
// It panic if error.
func MustFromYaml(s string) interface{} {
	d, err := FromYaml(s)
	if err != nil {
		panic(err)
	}
	return d
}

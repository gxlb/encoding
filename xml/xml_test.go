package xml

import (
	"fmt"
	"testing"
)

func TestXml(t *testing.T) {
	enc := Encoder{}
	var xx = []interface{}{
		123,
		"string",
		true,
		12.34,
		[]string{"a", "b", "c"},
		[]int{1, 2, 3, 4},
		map[string]interface{}{"a": 1, "b": "xx", "c": true},
		struct {
			X int
			Y string
		}{1, "yy"},
	}
	b, err := enc.Encode(xx, nil, true)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("xml:\n%s\n", string(b))

	dec := Decoder{}
	dec.Decode(nil, b)

}

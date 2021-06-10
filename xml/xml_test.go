package xml

import (
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
	}
	b, err := enc.Encode(xx, nil)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("xml:\n%s\n", string(b))
}

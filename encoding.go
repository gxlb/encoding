package encoding

import (
	"reflect"
	"unsafe"
)

//-------------------------

// unsafeByteString convert []byte to string without copy
// the origin []byte **MUST NOT** accessed after that
func unsafeByteString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// unsafeStringBytes return GoString's buffer slice
// ** NEVER modify returned []byte **
func unsafeStringBytes(s string) []byte {
	var bh reflect.SliceHeader
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return *(*[]byte)(unsafe.Pointer(&bh))
}

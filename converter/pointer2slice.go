package converter

import (
	"unsafe"
	"reflect"
)

//PointerToByteSlice 将unsafe.Pointer转成[]byte
func PointerToByteSlice(in unsafe.Pointer, len int) []byte {
	out := make([]byte, len)
	p := uintptr(in)
	for i := range out {
		out[i] = *(*byte)(unsafe.Pointer(p))
		p++
	}
	return out
}

//PointerToByteSliceWithoutCopy 将unsafe.Pointer转成[]byte 不拷贝内容
func PointerToByteSliceWithoutCopy(in unsafe.Pointer, len int) []byte {
	var out []byte
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&out)))
	sliceHeader.Cap = len
	sliceHeader.Len = len
	sliceHeader.Data = uintptr(in)
	return out
}

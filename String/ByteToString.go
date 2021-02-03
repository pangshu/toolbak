package String

import "unsafe"

// FromBytes converts the specified byte array to a string.
func (*String)ByteToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

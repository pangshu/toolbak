package Convert

import "unsafe"

// byte转string
func (*Convert)FromBytes(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

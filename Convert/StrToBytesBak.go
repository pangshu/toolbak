package Convert

//import (
//	"reflect"
//	"unsafe"
//)
//
//// StrToBytes 将字符串转换为字节切片.
//// 该方法零拷贝,但不安全.它直接转换底层指针,两者指向的相同的内存,改一个另外一个也会变.
//// 仅当临时需将长字符串转换且不长时间保存时可以使用.
//// 转换之后若没做其他操作直接改变里面的字符,则程序会崩溃.
//// 如 b:=String2bytes("xxx"); b[1]='d'; 程序将panic.
//func StrToBytes(val string) []byte {
//	pSliceHeader := &reflect.SliceHeader{}
//	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&val))
//	pSliceHeader.Data = strHeader.Data
//	pSliceHeader.Len = strHeader.Len
//	pSliceHeader.Cap = strHeader.Len
//	return *(*[]byte)(unsafe.Pointer(pSliceHeader))
//}
//
//func StringToBytes(str string) []byte {
//	x := (*[2]uintptr)(unsafe.Pointer(&str))
//	h := [3]uintptr{x[0], x[1], x[1]}
//	return *(*[]byte)(unsafe.Pointer(&h))
//}

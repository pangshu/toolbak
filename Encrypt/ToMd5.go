package Encrypt

import (
	"crypto/md5"
	"fmt"
)

// md5Byte 计算字节切片的 MD5 散列值.
func ToMd5(str string, lengthType ...int) string {
	var res string
	var length = 32
	if len(lengthType) > 0 {
		length = lengthType[0]
	}

	data := []byte(str)
	has := md5.Sum(data)
	dst := fmt.Sprintf("%x", has)

	if length == 16 {
		res = dst[8:24]
	} else {
		res = dst
	}

	return res
}

//// md5Byte 计算字节切片的 MD5 散列值.
//func md5Byte(str []byte, length uint8) []byte {
//	var res []byte
//	h := md5.New()
//	h.Write(str)
//
//	hBytes := h.Sum(nil)
//	dst := make([]byte, hex.EncodedLen(len(hBytes)))
//	hex.Encode(dst, hBytes)
//	if length == 16 {
//		res = dst[8:24]
//	} else {
//		res = dst
//	}
//
//	return res
//}
package Encrypt

// EasyEncrypt 简单加密.
// data为要加密的原字符串,key为密钥.
func EasyEncrypt(data, key []byte) []byte {
	dataLen := len(data)
	if dataLen == 0 {
		return nil
	}

	keyByte := ToMd5(string(key))
	keyLen := len(keyByte)

	var i, x, c int
	var res []byte
	for i = 0; i < dataLen; i++ {
		if x == keyLen {
			x = 0
		}

		c = (int(data[i]) + int(keyByte[x])) % 256
		res = append(res, byte(c))

		x++
	}

	res = append([]byte(keyByte[:len(key)]), Base64UrlEncode(res)...)
	return res
}

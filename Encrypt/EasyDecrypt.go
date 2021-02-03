package Encrypt

// EasyDecrypt 简单解密.
// val为待解密的字符串,key为密钥.
func EasyDecrypt(val, key []byte) []byte {
	keyLen := len(key)
	if len(val) <= keyLen {
		return nil
	}

	data, err := Base64UrlDecode(val[keyLen:])
	if err != nil {
		return nil
	}

	keyStr := ToMd5(string(key))
	if string(val[:keyLen]) != keyStr[:keyLen] {
		return nil
	}

	dataLen := len(data)
	keyStrLen := len(keyStr)

	var i, x, c int
	var res []byte
	for i = 0; i < dataLen; i++ {
		if x == keyStrLen {
			x = 0
		}

		if data[i] < keyStr[x] {
			c = int(data[i]) + 256 - int(keyStr[x])
		} else {
			c = int(data[i]) - int(keyStr[x])
		}
		res = append(res, byte(c))

		x++
	}

	return res
}

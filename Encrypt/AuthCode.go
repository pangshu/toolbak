package Encrypt

import (
	"fmt"
	"strconv"
	"time"
)

// AuthCode 授权码编码或解码;encode为true时编码,为false解码;expiry为有效期,秒;返回结果为加密/解密的字符串和有效期时间戳.
// key 的长度最大为32位
func AuthCode(str, key string, encode bool, expiry int64) ([]byte, int64) {
	// 加入随机密钥，可以令密文无任何规律，即便是原文和密钥完全相同，加密结果也会每次不同，增大破解难度。
	// 取值越大，密文变动规律越大，密文变化 = 16 的 DYNAMIC_KEY_LEN 次方
	// 当此值为 0 时，则不产生随机密钥

	keyLen := len(key)
	strLen := len(str)
	if str == "" || strLen == 0 {
		return nil, 0
	} else if !encode && strLen < keyLen {
		return nil, 0
	}

	// 密钥
	keyByte := ToMd5(key)

	// 密钥a会参与加解密
	keyA := keyByte[:16]

	// 密钥b会用来做数据完整性验证
	keyB := keyByte[16:]

	// 密钥c用于变化生成的密文
	var keyC string
	if encode == false {
		keyC = str[:keyLen]
	} else {
		now, _ := time.Now().MarshalBinary()
		tmpLen := 32 - keyLen
		timeBytes := ToMd5(string(now))
		keyC = timeBytes[tmpLen:]
	}

	// 参与运算的密钥
	keyD := ToMd5(string(append([]byte(keyA), []byte(keyC)...)))
	cryptKey := append([]byte(keyA), []byte(keyD)...)
	cryptKeyLen := len(cryptKey)
	// 明文，前10位用来保存时间戳，解密时验证数据有效性，10到26位用来保存密钥b，解密时会通过这个密钥验证数据完整性
	if encode == false { //解密
		//var err error
		strByte, err := Base64UrlDecode([]byte(str[keyLen:]))
		if err != nil {
			return nil, 0
		}
		str = string(strByte)
	} else {
		if expiry != 0 {
			expiry = expiry + time.Now().Unix()
		}
		expMd5 := ToMd5(string(append([]byte(str), []byte(keyB)...)), 16)
		str = fmt.Sprintf("%010d%s%s", expiry, expMd5, str)
		//str = append([]byte(fmt.Sprintf("%010d", expiry)), append(expMd5, str...)...)
	}

	strLen = len(str)
	resData := make([]byte, 0, strLen)
	var rndKey, box [256]int
	// 产生密钥簿
	h := 0
	i := 0
	j := 0

	for i = 0; i < 256; i++ {
		rndKey[i] = int(cryptKey[i % cryptKeyLen])
		box[i] = i
	}
	// 用固定的算法，打乱密钥簿，增加随机性，好像很复杂，实际上并不会增加密文的强度
	for i = 0; i < 256; i++ {
		j = (j + box[i] + rndKey[i]) % 256
		box[i], box[j] = box[j], box[i]
	}
	// 核心加解密部分
	h = 0
	j = 0
	for i = 0; i < strLen; i++ {
		h = ((h + 1) % 256)
		j = ((j + box[h]) % 256)
		box[h], box[j] = box[j], box[h]
		// 从密钥簿得出密钥进行异或，再转成字符
		resData = append(resData, byte(int(str[i])^box[(box[h]+box[j])%256]))
	}
	if encode == false { //解密
		// substr($result, 0, 10) == 0 验证数据有效性
		// substr($result, 0, 10) - time() > 0 验证数据有效性
		// substr($result, 10, 16) == substr(md5(substr($result, 26).$keyb), 0, 16) 验证数据完整性
		// 验证数据有效性，请看未加密明文的格式
		if len(resData) <= 26 {
			return nil, 0
		}

		expTime, _ := strconv.ParseInt(string(resData[:10]), 10, 0)
		if (expTime == 0 || expTime-time.Now().Unix() > 0) && string(resData[10:26]) == string(ToMd5(string(append(resData[26:], []byte(keyB)...)), 16)) {
			return resData[26:], expTime
		} else {
			return nil, expTime
		}
	} else { //加密
		// 把动态密钥保存在密文里，这也是为什么同样的明文，生产不同密文后能解密的原因
		resData = append([]byte(keyC), Base64UrlEncode(resData)...)
		return resData, expiry
	}
}

package Encrypt

// AesCBCDecrypt AES-CBC密码分组链接(Cipher-block chaining)模式解密.
// cipherText为密文;key为密钥,长16/24/32;paddingType为填充方式,枚举(PKCS_NONE,PKCS_ZERO,PKCS_SEVEN),默认PKCS_SEVEN.
func AesCBCDecrypt(cipherText, key []byte, paddingType ...int) ([]byte, error) {
	return AesDecrypt(cipherText, key, "CBC", paddingType...)
}

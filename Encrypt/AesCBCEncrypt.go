package Encrypt

// AesCBCEncrypt AES-CBC密码分组链接(Cipher-block chaining)模式加密.加密无法并行,不适合对流数据加密.
// clearText为明文;key为密钥,长16/24/32;paddingType为填充方式,枚举(PKCS_ZERO,PKCS_SEVEN),默认PKCS_SEVEN.
func AesCBCEncrypt(clearText, key []byte, paddingType ...int) ([]byte, error) {
	return AesEncrypt(clearText, key, "CBC", paddingType...)
}

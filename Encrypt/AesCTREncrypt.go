package Encrypt

// AesECBEncrypt AES-CTR计算器(Counter)模式加密.
// clearText为明文;key为密钥,长16/24/32.
func AesCTREncrypt(clearText, key []byte) ([]byte, error) {
	return AesEncrypt(clearText, key, "CTR", -1)
}

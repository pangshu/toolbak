package Encrypt

// AesOFBEncrypt AES-OFB输出反馈(Output feedback)模式加密.适合对流数据加密.
// clearText为明文;key为密钥,长16/24/32.
func AesOFBEncrypt(clearText, key []byte) ([]byte, error) {
	return AesEncrypt(clearText, key, "OFB", -1)
}

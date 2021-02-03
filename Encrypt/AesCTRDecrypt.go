package Encrypt
// AesECBDecrypt AES-CTR计算器(Counter)模式解密.
// cipherText为密文;key为密钥,长16/24/32.
func AesCTRDecrypt(cipherText, key []byte) ([]byte, error) {
	return AesDecrypt(cipherText, key, "CTR", -1)
}
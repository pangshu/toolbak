package Encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

const (
	DecPkcsNone = -1
	DecPkcsZero = 0
	DecPkcsSeven = 7
)
// AesDecrypt AES解密.
// cipherText为密文;key为密钥,长度16/24/32;
// mode为模式,枚举值(CBC,CFB,CTR,OFB);
// paddingType为填充方式,枚举(PKCS_NONE,PKCS_ZERO,PKCS_SEVEN),默认PKCS_SEVEN.
func AesDecrypt(cipherText, key []byte, mode string, paddingType ...int) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	pt := DecPkcsSeven
	if len(paddingType) > 0 {
		pt = paddingType[0]
	}

	blockSize := block.BlockSize()
	clan := len(cipherText)
	if clan < blockSize {
		return nil, errors.New("cipherText too short")
	}

	iv := key[:blockSize]
	origData := make([]byte, clan)

	switch mode {
	case "CBC":
		cipher.NewCBCDecrypter(block, iv).CryptBlocks(origData, cipherText)
	case "CFB":
		cipher.NewCFBDecrypter(block, iv).XORKeyStream(origData, cipherText)
	case "CTR":
		cipher.NewCTR(block, iv).XORKeyStream(origData, cipherText)
	case "OFB":
		cipher.NewOFB(block, iv).XORKeyStream(origData, cipherText)
	}

	if pt != DecPkcsNone && clan > 0 && int(cipherText[clan-1]) > clan {
		return nil, errors.New(fmt.Sprintf("aes [%s] decrypt failed", mode))
	}

	var plainText []byte
	switch pt {
	case DecPkcsZero:
		plainText = zeroUnPadding(cipherText)
	case DecPkcsSeven:
		plainText = pkcs7UnPadding(cipherText, blockSize)
	case DecPkcsNone:
		plainText = cipherText
	}

	return plainText, nil
}

package Encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

const (
	EncPkcsNone = -1
	EncPkcsZero = 0
	EncPkcsSeven = 7
)

// aesEncrypt AES加密.
// clearText为明文;key为密钥,长度16/24/32;
// mode为模式,枚举值(CBC,CFB,CTR,OFB);
// paddingType为填充方式,枚举(PKCS_NONE,PKCS_ZERO,PKCS_SEVEN),默认PKCS_SEVEN.
func AesEncrypt(clearText, key []byte, mode string, paddingType ...int) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	pt := EncPkcsSeven
	blockSize := block.BlockSize()
	if len(paddingType) > 0 {
		pt = paddingType[0]
	}
	switch pt {
	case EncPkcsZero:
		clearText = zeroPadding(clearText, blockSize)
	case EncPkcsSeven:
		clearText = pkcs7Padding(clearText, blockSize, false)
	}

	cipherText := make([]byte, len(clearText))
	//初始化向量
	iv := key[:blockSize]
	_, _ = io.ReadFull(rand.Reader, iv)
	//if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	//	return nil, err
	//}

	switch mode {
	case "CBC":
		cipher.NewCBCEncrypter(block, iv).CryptBlocks(cipherText, clearText)
	case "CFB":
		cipher.NewCFBEncrypter(block, iv).XORKeyStream(cipherText, clearText)
	case "CTR":
		cipher.NewCTR(block, iv).XORKeyStream(cipherText, clearText)
	case "OFB":
		cipher.NewOFB(block, iv).XORKeyStream(cipherText, clearText)
	}

	return cipherText, nil
}

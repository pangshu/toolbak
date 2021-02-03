package Encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// RsaPublicEncrypt RSA公钥加密.
// clearText为明文,publicKey为公钥.
func RsaPublicEncrypt(clearText, publicKey []byte) ([]byte, error) {
	// 解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 类型断言
	pubKey := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pubKey, clearText)
}


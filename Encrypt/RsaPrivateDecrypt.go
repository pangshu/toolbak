package Encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// RsaPrivateDecrypt RSA私钥解密.比加密耗时.
// cipherText为密文,privateKey为私钥.
func RsaPrivateDecrypt(cipherText, privateKey []byte) ([]byte, error) {
	// 获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}

	// 解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipherText)
}
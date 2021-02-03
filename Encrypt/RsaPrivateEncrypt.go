package Encrypt

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// RsaPrivateEncrypt RSA私钥加密.比解密耗时.
// clearText为明文,privateKey为私钥.
func RsaPrivateEncrypt(clearText, privateKey []byte) ([]byte, error) {
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

	return rsa.SignPKCS1v15(nil, priv, crypto.Hash(0), clearText)
}

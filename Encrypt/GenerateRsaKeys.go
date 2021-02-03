package Encrypt

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// GenerateRsaKeys 生成RSA密钥对.bits为密钥位数,通常为1024或2048.
func GenerateRsaKeys(bits int) (private []byte, public []byte, err error) {
	// 生成私钥文件
	var privateKey *rsa.PrivateKey
	privateKey, err = rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	privateBuff := new(bytes.Buffer)
	_ = pem.Encode(privateBuff, block)

	// 生成公钥文件
	var derPkix []byte
	publicKey := &privateKey.PublicKey
	derPkix, _ = x509.MarshalPKIXPublicKey(publicKey)
	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derPkix,
	}
	publicBuff := new(bytes.Buffer)
	_ = pem.Encode(publicBuff, block)

	private = privateBuff.Bytes()
	public = publicBuff.Bytes()

	return
}


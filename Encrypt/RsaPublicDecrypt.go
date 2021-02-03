package Encrypt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"math/big"
)

// RsaPublicDecrypt RSA公钥解密.
// cipherText为密文,publicKey为公钥.
func RsaPublicDecrypt(cipherText, publicKey []byte) ([]byte, error) {
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

	c := new(big.Int)
	m := new(big.Int)
	m.SetBytes(cipherText)
	e := big.NewInt(int64(pubKey.E))
	c.Exp(m, e, pubKey.N)
	out := c.Bytes()
	olen := len(out)
	skip := 0
	for i := 2; i < olen; i++ {
		if (i+1 < olen) && out[i] == 0xff && out[i+1] == 0 {
			skip = i + 2
			break
		}
	}

	return out[skip:], nil
}

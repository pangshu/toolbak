package Encrypt

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

// HmacShaX HmacSHA-x加密,x为1/256/512 .
func HmacShaX(data, secret []byte, x uint16) []byte {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	var h hash.Hash
	switch x {
	case 1:
		h = hmac.New(sha1.New, secret)
		break
	case 256:
		h = hmac.New(sha256.New, secret)
		break
	case 512:
		h = hmac.New(sha512.New, secret)
		break
	default:
		panic("[HmacShaX] x must be in [1, 256, 512]")
	}

	// Write Data to it
	h.Write(data)

	src := h.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	return dst
}

package Encrypt

// pkcs7UnPadding PKCS7拆解.
// origData为源数据;blockSize为分组长度.
func pkcs7UnPadding(origData []byte, blockSize int) []byte {
	olen := len(origData)
	if origData == nil || olen == 0 || blockSize <= 0 || olen%blockSize != 0 {
		return nil
	}

	unpadding := int(origData[olen-1])
	if unpadding == 0 || unpadding > olen {
		return nil
	}

	return origData[:(olen - unpadding)]
}

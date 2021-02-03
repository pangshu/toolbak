package Encrypt

import "golang.org/x/crypto/bcrypt"

// PasswordVerify 验证密码是否和散列值匹配.
func PasswordVerify(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

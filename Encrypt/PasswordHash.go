package Encrypt

import "golang.org/x/crypto/bcrypt"

// PasswordHash 创建密码的散列值;costs为算法的cost,范围4~31,默认10;注意:值越大越耗时.
func PasswordHash(password []byte, costs ...int) ([]byte, error) {
	var cost int
	if len(costs) == 0 {
		cost = 10
	} else {
		cost = costs[0]
		if cost < 4 {
			cost = 4
		} else if cost > 31 {
			cost = 15
		}
	}

	res, err := bcrypt.GenerateFromPassword(password, cost)
	return res, err
}

package Os

import "os"

func (*Os)GetPwd() (string, error) {
	dir, err := os.Getwd()
	return dir, err
}

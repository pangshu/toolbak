package Convert

import "strconv"

// BinToHex 将二进制字符串转换为十六进制字符串.
func (*Convert)BinToHex(str string) (string, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 16), nil
}

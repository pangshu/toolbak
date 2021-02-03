package Convert

import "strconv"

// HexToBin 将十六进制字符串转换为二进制字符串.
func (*Convert)HexToBin(data string) (string, error) {
	i, err := strconv.ParseInt(data, 16, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 2), nil
}

package Convert

import "strconv"

// BinToDec 将二进制转换为十进制.
func (*Convert)BinToDec(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return 0, err
	}
	return i, nil
}

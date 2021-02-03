package Convert

import "strconv"

// BaseConvert 进制转换,在任意进制之间转换数字.
// number为输入数值,fromNum为原进制,toNum为结果进制.
func (*Convert)BaseConvert(number string, fromNum, toNum int) (string, error) {
	i, err := strconv.ParseInt(number, fromNum, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toNum), nil
}

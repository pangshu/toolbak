package Regexp

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

// IsEmail 检查字符串是否邮箱.参数validateTrue,是否验证邮箱主机的真实性.
func IsEmail(email string, validateHost bool) (bool, error) {
	//长度检查
	length := len(email)
	at := strings.LastIndexByte(email, '@')
	if (length < 6 || length > 254) || (at <= 0 || at > length-3) {
		return false, fmt.Errorf("invalid email length")
	}

	//验证邮箱格式
	chkFormat := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString(email)
	if !chkFormat {
		return false, fmt.Errorf("invalid email format")
	}

	//验证主机
	if validateHost {
		host := email[at+1:]
		if _, err := net.LookupMX(host); err != nil {
			//因无法确定mx主机的smtp端口,所以去掉Hello/Mail/Rcpt检查邮箱是否存在
			//仅检查主机是否有效
			//TODO 仅对国内几家大的邮件厂家进行检查
			if _, err := net.LookupIP(host); err != nil {
				return false, err
			}
		}
	}

	return true, nil
}


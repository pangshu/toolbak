package Regexp

import "regexp"

// IsTel 是否固定电话或400/800电话.
func IsTel(str string) bool {
	return str != "" && regexp.MustCompile(`(` + `^(010|02\d{1}|0[3-9]\d{2})-\d{7,9}(-\d+)?$` + `)|(` + `^[48]00\d?(-?\d{3,4}){2}$` + `)`).MatchString(str)
}

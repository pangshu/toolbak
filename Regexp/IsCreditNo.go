package Regexp
//
//import (
//	"regexp"
//	"strconv"
//	"strings"
//	"github.com/pangshu/toolbak/Date"
//	"github.com/pangshu/toolbak/String"
//)
//
//// IsCreditNo 检查是否(15或18位)身份证号码,并返回经校验的号码.
//func IsCreditNo(str string) (bool, string) {
//	// 身份证区域
//	var CreditArea = map[string]string{"11": "北京", "12": "天津", "13": "河北", "14": "山西", "15": "内蒙古", "21": "辽宁", "22": "吉林", "23": "黑龙江", " 31": "上海", "32": "江苏", "33": "浙江", "34": "安徽", "35": "福建", "36": "江西", "37": "山东", "41": "河南", "42": "湖北", "43": "湖南", "44": "广东", "45": "广西", "46": "海南", "50": "重庆", "51": "四川", "52": "贵州", "53": "云南", "54": "西藏", "61": "陕西", "62": "甘肃", "63": "青海", "64": "宁夏", "65": "新疆", "71": "台湾", "81": "香港", "82": "澳门", "91": "国外"}
//
//	chk := str != "" && regexp.MustCompile(`(^[1-9]\d{5}(18|19|20)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$)`).MatchString(str)
//	if !chk {
//		return false, ""
//	}
//
//	// 检查省份代码
//	if _, chk = CreditArea[str[0:2]]; !chk {
//		return false, ""
//	}
//
//	// 将15位身份证升级到18位
//	leng := len(str)
//	if leng == 15 {
//		// 先转为17位,如果身份证顺序码是996 997 998 999,这些是为百岁以上老人的特殊编码
//		if chk, _ = String.Dstrpos(str[12:], []string{"996", "997", "998", "999"}, false); chk {
//			str = str[0:6] + "18" + str[6:]
//		} else {
//			str = str[0:6] + "19" + str[6:]
//		}
//
//		// 再加上校验码
//		code := append([]byte{}, creditChecksum(str))
//		str += string(code)
//	}
//
//	// 检查生日
//	birthday := str[6:10] + "-" + str[10:12] + "-" + str[12:14]
//	chk, tim := Date.IsDateToTime(birthday)
//	now := Date.UnixTime()
//	if !chk {
//		return false, ""
//	} else if tim >= now {
//		return false, ""
//	}
//
//	// 18位身份证需要验证最后一位校验位
//	if leng == 18 {
//		str = strings.ToUpper(str)
//		if str[17] != creditChecksum(str) {
//			return false, ""
//		}
//	}
//
//	return true, str
//}
//
//// creditChecksum 计算身份证校验码,其中id为身份证号码.
//func creditChecksum(id string) byte {
//	//∑(ai×Wi)(mod 11)
//	// 加权因子
//	factor := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
//	// 校验位对应值
//	code := []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
//
//	leng := len(id)
//	sum := 0
//	for i, char := range id[:leng-1] {
//		num, _ := strconv.Atoi(string(char))
//		sum += num * factor[i]
//	}
//
//	return code[sum%11]
//}
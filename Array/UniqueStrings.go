package Array

import "sort"

// UniqueStrings 移除字符串数组中的重复值.
func (*Array)UniqueStrings(strs []string) (res []string) {
	sort.Strings(strs)
	var last string
	for _, str := range strs {
		if str != last {
			res = append(res, str)
		}
		last = str
	}

	return
}

package Number

import (
	"toolbak/Array"
)

// IsNaturalRange 是否连续的自然数数组/切片,如[0,1,2,3...],其中不能有间断.
// strict为是否严格检查元素的顺序.
func (toolNumber *Number)IsNaturalRange(arr []int, strict bool) (res bool) {
	var toolArray Array.Array
	n := len(arr)
	if n == 0 {
		return
	}

	orig := toolNumber.Range(0, n-1)
	ctyp := "VALUE"

	if strict {
		ctyp = "ALL"
	}

	diff := toolArray.ArrayDiff(orig, arr, ctyp)

	res = len(diff) == 0
	return
}

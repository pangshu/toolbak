package Array

import (
	"reflect"
	"toolbak/Convert"
)

const (
	DiffTypeAll		string = "ALL"
	DiffTypeKey		string = "KEY"
	DiffTypeValue	string = "VALUE"
)
// ArrayDiff 计算数组(数组/切片/字典)的交集,返回在 arr1 中但不在 arr2 里的元素,注意会同时返回键.
// defaultType为两个数组的比较方式,枚举类型,有:
// VALUE 根据元素值比较, 返回在 arr1 中但是不在arr2 里的值;
// KEY 根据 arr1 中的键名和 arr2 进行比较,返回不同键名的项;
// ALL 同时比较键和值.
func (*Array)ArrayDiff(arr1, arr2 interface{}, defaultType ...string) map[interface{}]interface{} {
	var toolConvert Convert.Convert
	var compareType = DiffTypeAll
	if len(defaultType) > 0 {
		compareType = defaultType[0]
	}
	valA := reflect.ValueOf(arr1)
	valB := reflect.ValueOf(arr2)
	typA := valA.Kind()
	typB := valB.Kind()

	if (typA != reflect.Array && typA != reflect.Slice && typA != reflect.Map) || (typB != reflect.Array && typB != reflect.Slice && typB != reflect.Map) {
		panic("[ArrayDiff] arr1, arr2 type must be array, slice or map")
	}
	lenA := valA.Len()
	lenB := valB.Len()
	if lenA == 0 {
		return nil
	}

	resMap := make(map[interface{}]interface{})
	var iteA interface{}
	var chkKey bool
	var chkVal bool
	var chkRes bool

	if (typA == reflect.Array || typA == reflect.Slice) && (typB == reflect.Array || typB == reflect.Slice) {
		//两者都是数组/切片
		for i := 0; i < lenA; i++ {
			iteA = valA.Index(i).Interface()
			chkRes = true

			if compareType == DiffTypeAll {
				if i < lenB {
					chkRes = !reflect.DeepEqual(iteA, valB.Index(i).Interface())
				}
			} else if compareType == DiffTypeKey {
				chkRes = lenB > 0 && i >= lenB
			} else if compareType == DiffTypeValue {
				for j := 0; j < lenB; j++ {
					chkRes = !reflect.DeepEqual(iteA, valB.Index(j).Interface())
					if !chkRes {
						break
					}
				}
			}

			if chkRes {
				resMap[i] = iteA
			}
		}
	} else if (typA == reflect.Array || typA == reflect.Slice) && (typB == reflect.Map) {
		//A是数组/切片,B是字典
		for i := 0; i < lenA; i++ {
			iteA = valA.Index(i).Interface()
			chkRes = true

			for _, k := range valB.MapKeys() {
				chkKey = toolConvert.IsInt(k.Interface()) && toolConvert.ToInt(k.Interface()) == i
				chkVal = reflect.DeepEqual(iteA, valB.MapIndex(k).Interface())

				if compareType == DiffTypeKey && chkKey {
					chkRes = false
					break
				} else if compareType == DiffTypeValue && chkVal {
					chkRes = false
					break
				} else if compareType == DiffTypeAll && (chkKey && chkVal) {
					chkRes = false
					break
				}
			}

			if chkRes {
				resMap[i] = iteA
			}
		}
	} else if (typA == reflect.Map) && (typB == reflect.Array || typB == reflect.Slice) {
		//A是字典,B是数组/切片
		var kv int
		for _, k := range valA.MapKeys() {
			iteA = valA.MapIndex(k).Interface()
			chkRes = true

			if toolConvert.IsInt(k.Interface()) {
				kv = toolConvert.ToInt(k.Interface())
			} else {
				kv = -1
			}

			if compareType == DiffTypeAll {
				if kv >= 0 && kv < lenB {
					chkRes = !reflect.DeepEqual(iteA, valB.Index(kv).Interface())
				}
			} else if compareType == DiffTypeKey {
				chkRes = (kv < 0 || kv >= lenB)
			} else if compareType == DiffTypeValue {
				for i := 0; i < lenB; i++ {
					chkRes = !reflect.DeepEqual(iteA, valB.Index(i).Interface())
					if !chkRes {
						break
					}
				}
			}

			if chkRes {
				resMap[k.Interface()] = iteA
			}
		}
	} else if (typA == reflect.Map) && (typB == reflect.Map) {
		//两者都是字典
		var kv string
		for _, k := range valA.MapKeys() {
			iteA = valA.MapIndex(k).Interface()
			chkRes = true
			kv = toolConvert.ToStr(k.Interface())

			for _, k2 := range valB.MapKeys() {
				chkKey = kv == toolConvert.ToStr(k2.Interface())
				chkVal = reflect.DeepEqual(iteA, valB.MapIndex(k2).Interface())

				if compareType == DiffTypeKey && chkKey {
					chkRes = false
					break
				} else if compareType == DiffTypeValue && chkVal {
					chkRes = false
					break
				} else if compareType == DiffTypeAll && (chkKey || chkVal) {
					chkRes = false
					break
				}
			}

			if chkRes {
				resMap[k.Interface()] = iteA
			}
		}
	}

	return resMap
}
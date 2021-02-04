package Array

import (
	"reflect"
	"github.com/pangshu/toolbak/Convert"
)

const (
	IntersectTypeAll		string = "ALL"
	IntersectTypeKey		string = "KEY"
	IntersectTypeValue		string = "VALUE"
)
// ArrayIntersect 计算数组(数组/切片/字典)的交集,返回在 arr1 中且在 arr2 里的元素,注意会同时返回键.
// compareType为两个数组的比较方式,枚举类型,有:
// COMPARE_ONLY_VALUE 根据元素值比较, 返回在 arr1 中且在arr2 里的值;
// COMPARE_ONLY_KEY 根据 arr1 中的键名和 arr2 进行比较,返回相同键名的项;
// COMPARE_BOTH_KEYVALUE 同时比较键和值.
func (*Array)ArrayIntersect(arr1, arr2 interface{}, defaultType ...string) map[interface{}]interface{} {
	var toolConvert Convert.Convert
	var compareType = IntersectTypeAll
	if len(defaultType) > 0 {
		compareType = defaultType[0]
	}
	valA := reflect.ValueOf(arr1)
	valB := reflect.ValueOf(arr2)
	typA := valA.Kind()
	typB := valB.Kind()

	if (typA != reflect.Array && typA != reflect.Slice && typA != reflect.Map) || (typB != reflect.Array && typB != reflect.Slice && typB != reflect.Map) {
		panic("[ArrayIntersect] arr1, arr2 type must be array, slice or map")
	}

	lenA := valA.Len()
	lenB := valB.Len()
	if lenA == 0 || lenB == 0 {
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
			chkRes = false

			if compareType == IntersectTypeAll {
				if i < lenB {
					chkRes = reflect.DeepEqual(iteA, valB.Index(i).Interface())
				}
			} else if compareType == IntersectTypeKey {
				chkRes = i < lenB
			} else if compareType == IntersectTypeValue {
				for j := 0; j < lenB; j++ {
					chkRes = reflect.DeepEqual(iteA, valB.Index(j).Interface())
					if chkRes {
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
			chkRes = false

			for _, k := range valB.MapKeys() {
				chkKey = toolConvert.IsInt(k.Interface()) && toolConvert.ToInt(k.Interface()) == i
				chkVal = reflect.DeepEqual(iteA, valB.MapIndex(k).Interface())

				if compareType == IntersectTypeKey && chkKey {
					chkRes = true
					break
				} else if compareType == IntersectTypeValue && chkVal {
					chkRes = true
					break
				} else if compareType == IntersectTypeAll && (chkKey && chkVal) {
					chkRes = true
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
			chkRes = false

			if toolConvert.IsInt(k.Interface()) {
				kv = toolConvert.ToInt(k.Interface())
			} else {
				kv = -1
			}

			if compareType == IntersectTypeAll {
				if kv >= 0 && kv < lenB {
					chkRes = reflect.DeepEqual(iteA, valB.Index(kv).Interface())
				}
			} else if compareType == IntersectTypeKey {
				chkRes = kv >= 0 && kv < lenB
			} else if compareType == IntersectTypeValue {
				for i := 0; i < lenB; i++ {
					chkRes = reflect.DeepEqual(iteA, valB.Index(i).Interface())
					if chkRes {
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
			chkRes = false
			kv = toolConvert.ToStr(k.Interface())

			for _, k2 := range valB.MapKeys() {
				chkKey = kv == toolConvert.ToStr(k2.Interface())
				chkVal = reflect.DeepEqual(iteA, valB.MapIndex(k2).Interface())

				if compareType == IntersectTypeKey && chkKey {
					chkRes = true
					break
				} else if compareType == IntersectTypeValue && chkVal {
					chkRes = true
					break
				} else if compareType == IntersectTypeAll && (chkKey && chkVal) {
					chkRes = true
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

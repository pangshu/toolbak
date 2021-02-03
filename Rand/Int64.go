package Rand

import (
	"math/rand"
	"time"
)

//随机区间值
func (*Rand)Int64(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min+1)
	//MaxInt64 := int64(^uint(0) >> 1)
	//MinInt64 := ^MaxInt64
	//
	//if min > max {
	//	panic("[RandInt64]: min cannot be greater than max")
	//}
	//
	//if min == max {
	//	return min
	//}
	//
	////范围是否在边界内
	//mMax := int64(math.MaxInt32)
	//mMin := int64(math.MinInt32)
	//inrang := (mMin <= min && max <= mMax) || (MinInt64 <= min && max <= 0) || (0 <= min && max <= MaxInt64)
	//
	//if !inrang {
	//	panic("[RandInt64]: min and max exceed capacity,the result should be overflows int64.")
	//}
	//
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//return r.Int63n(max-min) + min
}

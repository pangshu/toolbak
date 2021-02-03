package Rand

import (
	"math"
	"math/rand"
	"time"
)

// Float64 生产一个随机float64整数.
func (*Rand)Float64(min, max float64) float64 {
	if min > max {
		panic("[RandFloat64]: min cannot be greater than max")
	}

	//范围是否在边界内
	mMax := float64(math.MaxFloat32)
	mMin := -mMax
	inrang := (mMin <= min && max <= mMax) || (-math.MaxFloat64 <= min && max <= 0) || (0 <= min && max <= math.MaxFloat64)
	if !inrang {
		panic("[RandFloat64]: min and max exceed capacity,the result should be overflows float64.")
	}

	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Seed(time.Now().UnixNano())
	num := rand.Float64()

	res := min + num*(max-min)
	return res
}

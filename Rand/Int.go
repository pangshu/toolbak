package Rand

import (
	"math/rand"
	"time"
)

func (*Rand)Int(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

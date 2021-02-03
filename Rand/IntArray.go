package Rand

import "math/rand"

// IntArray returns a random integer array with the specified start, end and size.
func (*Rand)IntArray(start, end, size int) []int {
	if end-start < size {
		size = end - start
	}

	var slice []int
	for i := start; i < end; i++ {
		slice = append(slice, i)
	}

	var ret []int
	for i := 0; i < size; i++ {
		idx := rand.Intn(len(slice))
		if slice != nil {
			ret = append(ret, slice[idx])
			slice = append(slice[:idx], slice[idx+1:]...)
		}

	}

	return ret
}

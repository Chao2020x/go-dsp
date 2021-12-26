package utils

import (
	"math/rand"
	"time"
)

func RandArray_Float64(length int) []float64 {
	rand.Seed(time.Now().UnixNano())
	var y = make([]float64, length)
	for i := 0; i < length; i++ {
		y[i] = rand.Float64()
	}
	return y
}

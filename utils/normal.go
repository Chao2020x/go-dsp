package utils

import (
	"math"

	"github.com/wcharczuk/go-chart/v2"
)

// Normal_Float64 映射到[-1,1]
func Normal_Float64(x []float64) []float64 {
	x_len := len(x)
	y := make([]float64, x_len)

	var (
		_Sum       = 0.0
		_Average   = 0.0
		_Min, _Max = chart.MinMax(x...)
	)
	for i := 0; i < x_len; i++ {
		_Sum += x[i]
	}
	_Average = _Sum / float64(x_len)
	_K := Max_Abs(_Max-_Average, _Min-_Average)
	for i := 0; i < x_len; i++ {
		y[i] = (x[i] - _Average) / _K
	}
	return y
}

func Max_Abs(x, y float64) float64 {
	var zmax float64 = 0
	if math.Abs(x) > math.Abs(y) {
		zmax = math.Abs(x)
	} else {
		zmax = math.Abs(y)
	}
	return zmax
}

package utils

//错了
func Conv(x, y []float64) []float64 {
	x_len := len(x)
	y_len := len(y)
	z_len := x_len + y_len - 1
	z := make([]float64, z_len)

	//通过离散卷积公式计算
	for i := 0; i < z_len; i++ {
		var k float64 = 0
		for j := 0; j < x_len; j++ {
			if (i-j) >= 0 && (i-j) < y_len {
				k += x[j] * y[i-j]
			}
		}
		z[i] = k
	}
	return z
}

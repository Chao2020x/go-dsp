package utils

import (
	"github.com/chfenger/goNum"
)

// Enframe 分帧操作，相当于将信号分解为若干个信号片段，并将片段与窗函数进行对应元素的乘法
// 分帧就是通过加窗函数实现的，假设原始信号为 _X(n)，窗函数为 _Win(n) 分帧就是 _X(n) * _Win(n)
// _Inc 分帧通常有一定的交叠部分，就是帧移
func Enframe(_X []float64, _Win []float64, _Inc int) *goNum.Matrix {
	_Nlen := len(_Win)
	_NX := len(_X)
	_Nf := (_NX - _Nlen + _Inc) / _Inc

	_Frameout := goNum.ZeroMatrix(_Nf, _Nlen)
	for i := 0; i < _Nf; i++ {
		for j := 0; j < _Nlen; j++ {
			_Frameout.SetMatrix(i, j, _X[_Inc*i+j])
		}
	}

	for i := 0; i < _Nf; i++ {
		for j := 0; j < _Nlen; j++ {
			var val float64 = _Frameout.GetFromMatrix(i, j) * _Win[j]
			_Frameout.SetMatrix(i, j, val)
		}
	}
	return &_Frameout
}

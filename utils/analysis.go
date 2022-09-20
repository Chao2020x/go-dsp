package utils

import (
	"math/cmplx"

	"go-dsp/utils/fft"

	"github.com/chfenger/goNum"
)

//短时时域分析

//计算短时能量函数
func ShortTime_Energy(_X, _Win []float64, _Inc int) []float64 {
	_EnXMat := Enframe(_X, _Win, _Inc)
	var _EnArray = make([]float64, _EnXMat.Rows)
	for i := range _EnArray {
		var _temp = _EnXMat.RowOfMatrix(i)
		_EnArray[i] = Sum(Multiply(_temp, _temp))
	}
	return _EnArray
}

//计算短时平均幅度
func ShortTime_AverageMagnitude(_X, _Win []float64, _Inc int) []float64 {
	_AmXMat := Enframe(_X, _Win, _Inc)
	var _AmArray = make([]float64, _AmXMat.Rows)
	for i := range _AmArray {
		_AmArray[i] = Average(Abs(_AmXMat.RowOfMatrix(i)))
	}
	return _AmArray
}

//短时过零率
func ShortTime_ZeroCR(_X, _Win []float64, _Inc int) []float64 {
	_ZcrMat := Enframe(_X, _Win, _Inc)
	var _ZcrArray = make([]float64, _ZcrMat.Rows)
	for i := range _ZcrArray {
		_xxx := _ZcrMat.RowOfMatrix(i)
		_xxx111 := _xxx[:len(_xxx)-1]
		_xxx222 := _xxx[1:]
		_ZcrArray[i] = Sum(where_returnSelf(Multiply(_xxx111, _xxx222), 0, 1, 0))
	}
	return _ZcrArray
}

func where_returnSelf(_Array []float64, _BValue, _Greater, _Less float64) []float64 {
	for i := range _Array {
		if _Array[i] < _BValue {
			_Array[i] = _Greater
			continue
		}
		_Array[i] = _Less
	}
	return _Array
}

//计算短时相关函数
func ShortTime_AC(_X, _Win []float64, _Inc int) []float64 {
	_ACMat := Enframe(_X, _Win, _Inc)
	_ACMat_T := _ACMat.Transpose()
	_Para := goNum.ZeroMatrix(_ACMat_T.Rows, _ACMat_T.Columns)
	for i := 0; i < _ACMat_T.Columns; i++ {
		_Col_Array := _ACMat_T.ColumnOfMatrix(i)
		for r := 0; r < _Para.Rows; r++ {
			_Para.SetMatrix(r, i, Sum(Multiply(_Col_Array, _Col_Array)))
		}
	}
	return _Para.Transpose().Data
}

// 计算短时幅度差
func ShortTime_Amdf(_X, _Win []float64, _Inc int) []float64 {
	_AmdfMat_T := Enframe(_X, _Win, _Inc).Transpose() //转置矩阵
	_fn := _AmdfMat_T.Columns
	_wlen := _AmdfMat_T.Rows
	_Para := goNum.ZeroMatrix(_AmdfMat_T.Rows, _AmdfMat_T.Rows)
	for i := 0; i < _fn; i++ {
		_UtempMap := _AmdfMat_T.ColumnOfMatrix(i)
		for k := 0; k < _wlen; k++ {
			_PSum := Sum(Abs(Sub(_UtempMap[k:], _UtempMap[:_wlen-k])))
			for j := 0; j < _wlen; j++ {
				_Para.SetMatrix(k, j, _PSum)
			}
		}
	}
	return _Para.Data
}

func FrameTimeC(frameNum, frameLen, inc int, fs int) []float64 {
	ll := make([]float64, frameNum)
	for i := range ll {
		ll[i] = (float64((i-1)*inc) + (float64(frameLen) / 2)) / float64(fs)
	}
	return ll
}

// 短时频域分析

// 短时傅里叶变换函数
func ShortTime_FFT(_X, _Win []float64, _Nfft, _Inc int) *goNum.Matrix {
	_XN_T := Enframe(_X, _Win, _Inc).Transpose()
	_XN_2DArray := make([][]float64, _XN_T.Rows)

	for i := 0; i < _XN_T.Rows; i++ {
		_XN_2DArray[i] = make([]float64, _XN_T.Columns)
		copy(_XN_2DArray[i], _XN_T.RowOfMatrix(i))
	}

	fft_result := fft.FFT2Real(_XN_2DArray)
	rows := len(fft_result)
	if rows == 0 {
		return nil
	}
	cols := len(fft_result[0])
	_NewMat := goNum.ZeroMatrix(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			_NewMat.SetMatrix(i, j, cmplx.Abs(fft_result[i][j]))
		}
	}

	return &_NewMat
}

func abs_() {

}

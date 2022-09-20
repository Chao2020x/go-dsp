package chapter2

import (
	"fmt"
	"go-dsp/utils"
)

//实验要求一：语音信号叠加
func Chapter2_2_1() {

	x, fs, err := utils.WavRead_f64("chapter2\\C2_2_y.wav")
	if err != nil {
		fmt.Println("reads the WAV failed, err:", err)
		return
	}
	x = utils.Normal_Float64(x)

	t_len := len(x)
	var t []float64 = make([]float64, t_len)
	for i := 0; i < t_len; i++ {
		t[i] = float64(i) / float64(fs)
	}

	y := utils.RandArray_Float64(len(x))
	y = utils.Normal_Float64(y)

	z_len := len(x)
	var z = make([]float64, z_len)
	for i := 0; i < z_len; i++ {
		z[i] = x[i] + y[i]
	}
	z = utils.Normal_Float64(z)

	utils.NewVChart(t, x, "时间/s", "归一化幅值", "(a)原始信号").SaveWavePicture("chapter2\\1、(a)原始信号")
	utils.NewVChart(t, y, "时间/s", "归一化幅值", "(a)随机序列").SaveWavePicture("chapter2\\1、(b)随机序列")
	utils.NewVChart(t, z, "时间/s", "归一化幅值", "(a)随机序列").SaveWavePicture("chapter2\\1、(c)线性叠加")
}

// 实验要求二：语音信号卷积
func Chapter2_2_2() {
	x, fs, err := utils.WavRead_f64("chapter2\\C2_2_y.wav")
	if err != nil {
		fmt.Println("reads the WAV failed, err:", err)
		return
	}
	x = utils.Normal_Float64(x)

	t_len := len(x)
	var t []float64 = make([]float64, t_len)
	for i := 0; i < t_len; i++ {
		t[i] = float64(i) / float64(fs)
	}

	y := utils.RandArray_Float64(len(x))
	y = utils.Normal_Float64(y)

	z := utils.Conv(x, y)
	z = utils.Normal_Float64(z)

	z_len := len(z)
	var t2 = make([]float64, z_len)
	for i := 0; i < z_len; i++ {
		t2[i] = float64(i) / float64(fs)
	}

	utils.NewVChart(t, x, "时间/s", "归一化幅值", "(a)原始信号").SaveWavePicture("chapter2\\2、(a)原始信号")
	utils.NewVChart(t, y, "时间/s", "归一化幅值", "(a)随机序列").SaveWavePicture("chapter2\\2、(b)随机序列")
	utils.NewVChart(t2, z, "时间/s", "归一化幅值", "(a)信号卷积").SaveWavePicture("chapter2\\2、(c)信号卷积")
}

// 实验要求三：语音信号采样频率变换
func Chapter2_2_3() {
	x, fs, err := utils.WavRead_f64("chapter2\\C2_2_y.wav")
	if err != nil {
		fmt.Println("reads the WAV failed, err:", err)
		return
	}
	x = utils.Normal_Float64(x)

	t_len := len(x)
	var t []float64 = make([]float64, t_len)
	for i := 0; i < t_len; i++ {
		t[i] = float64(i) / float64(fs)
	}
	utils.NewVChart(t, x, "时间/s", "归一化幅值", "(a)原始信号").SaveWavePicture("chapter2\\3、(a)原始信号")
}

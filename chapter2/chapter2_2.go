package chapter2

import (
	"Sound_Signal_Process_Golang/utils"
	"fmt"
	"math"

	"github.com/wcharczuk/go-chart/v2"
)

//实验要求一：语音信号叠加
func Chapter2_2_1() {

	x, fs, err := utils.WavRead_f64("chapter2\\C2_2_y.wav")
	if err != nil {
		fmt.Println("reads the WAV failed, err:", err)
		return
	}

	var t []float64

	for i, _ := range x {
		t = append(t, float64(i)/float64(fs))
	}

	_min, _max := chart.MinMax(x...)

	xmax := 0.1
	if math.Abs(_max) > math.Abs(_min) {
		xmax = math.Abs(_max)
	} else {
		xmax = math.Abs(_min)
	}

	for i := 0; i < len(x); i++ {
		x[i] = x[i] / xmax
	}

	y := utils.RandArray_Float64(len(x))

	z_len := len(x)
	var z = make([]float64, z_len)

	for i := 0; i < z_len; i++ {
		z[i] = x[i] + y[i]
	}
	_min, _max = chart.MinMax(z...)
	zmax := 0.1
	if math.Abs(_max) > math.Abs(_min) {
		zmax = math.Abs(_max)
	} else {
		zmax = math.Abs(_min)
	}
	for i := 0; i < z_len; i++ {
		z[i] = z[i] / zmax
	}

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

	var t []float64

	for i, _ := range x {
		t = append(t, float64(i)/float64(fs))
	}

	_min, _max := chart.MinMax(x...)

	xmax := 0.1
	if math.Abs(_max) > math.Abs(_min) {
		xmax = math.Abs(_max)
	} else {
		xmax = math.Abs(_min)
	}

	for i := 0; i < len(x); i++ {
		x[i] = x[i] / xmax
	}

	y := utils.RandArray_Float64(len(x))

	z := utils.Conv(x, y)
	z_len := len(z)
	_min, _max = chart.MinMax(z...)
	zmax := 0.1
	if math.Abs(_max) > math.Abs(_min) {
		zmax = math.Abs(_max)
	} else {
		zmax = math.Abs(_min)
	}
	for i := 0; i < z_len; i++ {
		z[i] = z[i] / zmax
	}

	var t2 = make([]float64, z_len)
	for i := 0; i < z_len; i++ {
		t2[i] = float64(i) / float64(fs)
	}

	utils.NewVChart(t, x, "时间/s", "归一化幅值", "(a)原始信号").SaveWavePicture("chapter2\\2、(a)原始信号")
	utils.NewVChart(t, y, "时间/s", "归一化幅值", "(a)随机序列").SaveWavePicture("chapter2\\2、(b)随机序列")

	//TODO 这里为什么不对？？
	utils.NewVChart(t2, z, "时间/s", "归一化幅值", "(a)随机序列").SaveWavePicture("chapter2\\2、(c)信号卷积")
}

func Chapter2_2_3() {
	x, fs, err := utils.WavRead_f64("chapter2\\C2_2_y.wav")
	if err != nil {
		fmt.Println("reads the WAV failed, err:", err)
		return
	}

	var t []float64

	for i, _ := range x {
		t = append(t, float64(i)/float64(fs))
	}

	_min, _max := chart.MinMax(x...)

	xmax := 0.1
	if math.Abs(_max) > math.Abs(_min) {
		xmax = math.Abs(_max)
	} else {
		xmax = math.Abs(_min)
	}

	for i := 0; i < len(x); i++ {
		x[i] = x[i] / xmax
	}

	utils.NewVChart(t, x, "时间/s", "归一化幅值", "(a)原始信号").SaveWavePicture("chapter2\\3、(a)原始信号")
}

package chapter1

import (
	"Sound_Signal_Process_Golang/utils"
	"math"
	"math/cmplx"

	"github.com/mjibson/go-dsp/fft"
)

// ***************1.正弦波****************
func Example_1() {
	fs := 100 //采样频率
	N := 128
	var n []int
	for i := 0; i < N; i++ {
		n = append(n, i)
	}

	var t []float64
	for _, v := range n {
		t = append(t, float64(v)/float64(fs))
	}

	f0 := 10 //设定正弦信号频率

	//生成正弦信号
	var x []float64
	for _, value := range t {
		var s float64 = float64(2*f0) * value
		v2 := math.Sin(math.Pi * s)
		x = append(x, v2)
	}
	utils.NewVChart(t, x, "时间/s", "幅值", "时域波形").SaveWavePicture("chapter1\\正弦波-时域波形")

	fft_value := fft.FFTReal(x)
	var mag []float64
	for _, v := range fft_value {
		mag = append(mag, cmplx.Abs(v))
	}
	var f []float64
	for i, _ := range mag {
		kk := float64(i*fs) / float64(len(mag))
		f = append(f, kk)
	}
	utils.NewVChart(f, mag, "频率/Hz", "幅值", "幅频谱图").SaveWavePicture("chapter1\\正弦波-幅频谱图")

	var sq []float64
	for _, v := range fft_value {
		sq = append(sq, cmplx.Abs(v))
	}
	utils.NewVChart(f, sq, "频率/Hz", "均方根谱", "均方根谱").SaveWavePicture("chapter1\\正弦波-均方根谱")

	var power []float64
	for _, v := range sq {
		power = append(power, v*v)
	}
	utils.NewVChart(f, power, "频率/Hz", "功率谱", "功率谱").SaveWavePicture("chapter1\\正弦波-功率谱")

	var ln []float64
	for _, v := range sq {
		ln = append(ln, math.Log(v))
	}
	utils.NewVChart(f, ln, "频率/Hz", "对数谱", "对数谱").SaveWavePicture("chapter1\\正弦波-对数谱")

	var xifft []float64
	for _, v := range fft.IFFT(fft_value) {
		xifft = append(xifft, real(v))
	}

	var ti []float64
	for i, _ := range xifft {
		ti = append(ti, float64(i)/float64(fs))
	}
	utils.NewVChart(ti, xifft, "时间/s", "幅值", "IFFT 后的信号波形").SaveWavePicture("chapter1\\正弦波-IFFT后的信号波形")
}

// %****************2.白噪声****************%
func Example_2() {
	fs := 50 //设定采样频率

	var t []float64

	for i := 0; i < 100; i++ {
		t = append(t, float64(i)*0.1-5.0)
	}

	x := utils.RandArray_Float64(len(t))
	utils.NewVChart(t, x, "时间(s)", "幅值", "时域波形").SaveWavePicture("chapter1\\白噪声-时域波形")

	fft_value := fft.FFTReal(x)
	var mag []float64
	for _, v := range fft_value {
		mag = append(mag, cmplx.Abs(v))
	}
	var f []float64
	for i, _ := range mag {
		kk := float64(i*fs) / float64(len(mag))
		f = append(f, kk)
	}
	utils.NewVChart(f, mag, "频率/Hz", "幅值", "幅频谱图").SaveWavePicture("chapter1\\白噪声-幅频谱图")

	var sq []float64
	for _, v := range fft_value {
		sq = append(sq, cmplx.Abs(v))
	}
	utils.NewVChart(f, sq, "频率/Hz", "均方根谱", "均方根谱").SaveWavePicture("chapter1\\白噪声-均方根谱")

	var power []float64
	for _, v := range sq {
		power = append(power, v*v)
	}
	utils.NewVChart(f, power, "频率/Hz", "功率谱", "功率谱").SaveWavePicture("chapter1\\白噪声-功率谱")

	var ln []float64
	for _, v := range sq {
		ln = append(ln, math.Log(v))
	}
	utils.NewVChart(f, ln, "频率/Hz", "对数谱", "对数谱").SaveWavePicture("chapter1\\白噪声-对数谱")

	var xifft []float64
	for _, v := range fft.IFFT(fft_value) {
		xifft = append(xifft, real(v))
	}

	var ti []float64
	for i, _ := range xifft {
		ti = append(ti, float64(i)/float64(fs))
	}
	utils.NewVChart(ti, xifft, "时间/s", "幅值", "IFFT 后的信号波形").SaveWavePicture("chapter1\\白噪声-IFFT后的信号波形")
}

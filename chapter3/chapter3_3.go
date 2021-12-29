package chapter3

import (
	"Sound_Signal_Process_Golang/utils"
	"fmt"
	"math"

	"github.com/mjibson/go-dsp/window"
)

func Chapter3_3_1() {
	//读入文件
	x, fs, err := utils.WavRead_f64("chapter3\\C3_2_y.wav")
	if err != nil {
		fmt.Println("reads the WAV failed, err:", err)
		return
	}
	wlen, inc := 256, 128    //给出帧长和帧移
	win := window.Hann(wlen) //给出海宁窗
	nfft := wlen
	y_Mat := utils.ShortTime_FFT(x, win, nfft, inc) //求短时傅里叶变换

	fn := y_Mat.Columns

	//计算fft后频率刻度
	var freq = make([]float64, wlen/2)
	for i := range freq {
		freq[i] = float64(i*fs) / float64(wlen)
	}

	//计算每帧对应时间
	frametime := utils.FrameTimeC(fn, wlen, inc, fs)

	for i := range y_Mat.Data {
		y_Mat.Data[i] = y_Mat.Data[i] * y_Mat.Data[i]
	}
	copy(y_Mat.Data, utils.Reverse(y_Mat.Data)) //翻转
	for i := range y_Mat.Data {
		y_Mat.Data[i] = math.Log10(y_Mat.Data[i])
	}

	utils.NewImagesc("时间/s", "频率/Hz", "能量谱图", frametime, y_Mat).SaveWavePicture("chapter3\\3、能量谱图")

}

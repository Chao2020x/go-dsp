package chapter3

import (
	"fmt"
	"go-dsp/utils"

	"go-dsp/utils/window"
)

// Chapter3_2_1 短时时域分析
func Chapter3_2_1() {

	//读入文件
	x, fs, err := utils.WavRead_f64("chapter3\\C3_2_y.wav")
	if err != nil {
		fmt.Println("reads the WAV failed, err:", err)
		return
	}
	x = utils.Normal_Float64(x)

	wlen, inc := 200, 100    //给出帧长和帧移
	win := window.Hann(wlen) //给出海宁窗
	N := len(x)              //信号长度

	//计算出信号的时间刻度
	var time []float64 = make([]float64, N)
	for i := 0; i < N; i++ {
		time[i] = float64(i) / float64(fs)
	}

	En := utils.ShortTime_Energy(x, win, inc)           //短时能量
	Mn := utils.ShortTime_AverageMagnitude(x, win, inc) //短时平均幅度
	Zcr := utils.ShortTime_ZeroCR(x, win, inc)          //短时过零率
	Ac := utils.ShortTime_AC(x, win, inc)               //计算短时自相关
	Amdf := utils.ShortTime_Amdf(x, win, inc)           //计算短时幅度差

	//短时幅度差暂不写
	fn := len(En) //求出帧数
	frameTime := utils.FrameTimeC(fn, wlen, inc, fs)

	utils.NewVChart(time, x, "时间/s", "幅值", "语音波形").SaveWavePicture("chapter3\\2、(a)语音波形")
	utils.NewVChart(frameTime, Mn, "时间/s", "幅值", "短时幅度").SaveWavePicture("chapter3\\2、(b)短时幅度")
	utils.NewVChart(frameTime, En, "时间/s", "幅值", "短时能量").SaveWavePicture("chapter3\\2、(c)短时能量")
	utils.NewVChart(frameTime, Zcr, "时间/s", "幅值", "短时过零率").SaveWavePicture("chapter3\\2、(d)短时过零率")

	timeAC := make([]float64, len(Ac))
	for i := range timeAC {
		timeAC[i] = float64(i) / 10000.0
	}
	utils.NewVChart(timeAC, Ac, "点数x10^4", "幅值", "短时自相关").SaveWavePicture("chapter3\\2、(e)短时自相关")

	timeAmdf := make([]float64, len(Amdf))
	for i := range timeAmdf {
		timeAmdf[i] = float64(i) / 10000.0
	}
	utils.NewVChart(timeAmdf, Amdf, "点数x10^4", "幅值", "短时幅度差").SaveWavePicture("chapter3\\2、(f)短时幅度差")
}

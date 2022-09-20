package chapter2

import (
	"fmt"
	"go-dsp/utils"
)

//实验要求一：计算有效声压
func Chapter2_3_1() {
	x, fs, err := utils.WavRead_f64("chapter2\\C2_3_y.wav")
	if err != nil {
		fmt.Println("reads the WAV failed, err:", err)
		return
	}
	// x = utils.Normal_Float64(x)
	fmt.Println(len(x), fs, x[len(x)-3])

}

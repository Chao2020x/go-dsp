package utils

import (
	"fmt"
	"os"

	"github.com/go-audio/wav"
)

func Wav2Read_f32(Filename string) {
	file, err := os.OpenFile(Filename, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	gwav := wav.NewDecoder(file)

	gwav.ReadInfo()

	fmt.Println(gwav.SampleRate, gwav)
}

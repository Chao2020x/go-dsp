package utils

import (
	"fmt"
	"os"

	"go-dsp/utils/wav"
)

func WavRead_f64(Filename string) ([]float64, int, error) {

	data_f32, fs, err := WavRead_f32(Filename)
	var data_f64 = make([]float64, len(data_f32))
	if err != nil {
		fmt.Println("WavRead_f32 ReadFloats failed, err:", err)
		return nil, 0, err
	}
	_sample := len(data_f32)
	for i := 0; i < _sample; i++ {
		data_f64[i] = float64(data_f32[i])
	}
	return data_f64, fs, nil
}

func WavRead_f32(Filename string) ([]float32, int, error) {
	file, err := os.OpenFile(Filename, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return nil, 0, err
	}
	defer file.Close()
	gwav, err := wav.New(file)
	if err != nil {
		fmt.Println("New reads the WAV header failed, err:", err)
		return nil, 0, err
	}

	data_f32, err := gwav.ReadFloats(gwav.Samples)
	if err != nil {
		fmt.Println("WavRead_f32 ReadFloats failed, err:", err)
		return nil, 0, err
	}
	return data_f32, int(gwav.SampleRate), nil
}

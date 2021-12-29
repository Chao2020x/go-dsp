package chapter3

import (
	"fmt"

	"github.com/mjibson/go-dsp/window"
)

func Chapter3_1_1() {
	gg := window.Hann(9)
	fmt.Println(gg)
}

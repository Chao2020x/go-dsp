package chapter3

import (
	"fmt"

	"go-dsp/utils/window"
)

func Chapter3_1_1() {
	gg := window.Hann(9)
	fmt.Println(gg)
}

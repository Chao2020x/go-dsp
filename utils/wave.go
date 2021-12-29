package utils

import (
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/chfenger/goNum"
	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2"
)

type VChart struct {
	XAxisName string
	YAxisName string
	Title     string

	YValue []float64
	XValue []float64

	YValue_Min float64
	YValue_Max float64
}

func (v *VChart) MinMax() {
	v.YValue_Min, v.YValue_Max = chart.MinMax(v.YValue...)
}

func NewVChart(_XValue, _YValue []float64, _XAxisName, _YAxisName, _Title string) *VChart {
	e := &VChart{
		XAxisName: _XAxisName,
		YAxisName: _YAxisName,

		Title: _Title,

		YValue: _YValue,
		XValue: _XValue,
	}
	e.MinMax()
	return e
}

func (v *VChart) SaveWavePicture(_Filename string) {
	graph := chart.Chart{
		Title: v.Title,

		XAxis: chart.XAxis{
			Name: v.XAxisName,
			// ValueFormatter: chart.TimeMinuteValueFormatter, //TimeHourValueFormatter,
		},

		YAxis: chart.YAxis{
			Name:     v.YAxisName,
			AxisType: chart.YAxisPrimary,
			Range: &chart.ContinuousRange{
				Min: v.YValue_Min,
				Max: v.YValue_Max,
			},
		},

		Font: getZWFont(),

		Series: []chart.Series{
			chart.ContinuousSeries{ //TimeSeries{
				Name: _Filename,
				Style: chart.Style{
					StrokeColor: chart.GetDefaultColor(1),
				},
				YAxis:   chart.YAxisSecondary,
				XValues: v.XValue,
				YValues: v.YValue,
			},
		},
	}

	fileHandle, _ := os.Create(_Filename + ".png")
	defer fileHandle.Close()
	graph.Render(chart.PNG, fileHandle)
}

// getZWFont 加载字体
func getZWFont() *truetype.Font {

	fontFile := ".\\font\\Alibaba-PuHuiTi-Medium.ttf"

	// 读字体数据
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println(err)
		return nil
	}
	font, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Println(err)
		return nil
	}
	return font
}

type Imagesc struct {
	XAxisName string
	YAxisName string
	Title     string
	XRows     int
	YColumns  int

	Frametime []float64

	PointMat *goNum.Matrix
}

func NewImagesc(_XAxisName, _YAxisName, _Title string, _Frametime []float64, _Mat *goNum.Matrix) *Imagesc {

	return &Imagesc{
		XAxisName: _XAxisName,
		YAxisName: _YAxisName,
		Title:     _Title,
		XRows:     _Mat.Rows,
		YColumns:  _Mat.Columns,
		PointMat:  _Mat,
		Frametime: _Frametime,
	}
}

// SaveWavePicture 将结构体Imagesc中矩阵PointMat中的元素数值按大小转化为不同颜色
func (img *Imagesc) SaveWavePicture(_Filename string) {
	file, _ := os.Create(_Filename + ".png")
	defer file.Close()

	dx, dy := img.XRows, img.YColumns
	rgba := image.NewRGBA(image.Rect(0, 0, dx, dy))
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			rgba.Set(x, y, color.RGBA{uint8(x), uint8(y), 128, 255})
		}
	}
	if err := png.Encode(file, rgba); err != nil {
		panic(err)
	}
}

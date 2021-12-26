package utils

import (
	"io/ioutil"
	"log"
	"os"

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

	f, _ := os.Create(_Filename + ".png")
	defer f.Close()
	graph.Render(chart.PNG, f)

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

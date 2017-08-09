package main

import (
	"time"
	"bytes"
	"strconv"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func FloatToString(inputNum float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(inputNum, 'f', 0, 64)
}

func ShiftPoints(pts []float64){
	size := len(pts)
	for i := 0; i < size - 1; i++ {
		pts[i] = pts[i + 1]
	}
}

func InitArrays(cpuPoints []float64, ramPoints []float64, xAxisPoints []float64){
	for i, _ := range cpuPoints {
		cpuPoints[i] = 0
		ramPoints[i] = 0
		xAxisPoints[i] = float64(i * DELAY_SEC)
	}
}

func StartMonitoringResources(monChan chan *bytes.Buffer, monRestart chan bool){
	POINTS = (60 / DELAY_SEC) + 1

	var (
		cpuPoints = make([]float64, POINTS)
		ramPoints = make([]float64, POINTS)
		xAxisPoints = make([]float64, POINTS)
	)

	//init data arrays
	InitArrays(cpuPoints[:], ramPoints[:], xAxisPoints[:])

	for true {
		select {
			case <-monRestart:
				POINTS = (60 / DELAY_SEC) + 1

				cpuPoints = make([]float64, POINTS)
				ramPoints = make([]float64, POINTS)
				xAxisPoints = make([]float64, POINTS)

				InitArrays(cpuPoints[:], ramPoints[:], xAxisPoints[:])

			default:
		}

		//fixme no errors checks
		c, _ := cpu.Percent(DELAY, false) // get average % for  DELAY_SEC seconds
		m, _ := mem.VirtualMemory()

		cpuPoints[POINTS - 1] = c[0]
		ramPoints[POINTS - 1] = m.UsedPercent

		graph := chart.Chart{
			Width: 450,
			Height: 450,
			YAxis: chart.YAxis{
				Style: chart.Style{
					Show: true,
				},
				Range: &chart.ContinuousRange{
					Min: 0,
					Max: 100,
				},
				Ticks: []chart.Tick{
					{0.0, "0"},
					{50.0, "50"},
					{100.0, "100"},
				},
			},
			XAxis: chart.XAxis{
				Style: chart.Style{
					Show: true,
				},
				Range: &chart.ContinuousRange{
					Min: 0,
					Max: 60,
				},
				Ticks: []chart.Tick{
					{60.0, "0"},
					{50.0, "10"},
					{40.0, "20"},
					{30.0, "30"},
					{20.0, "40"},
					{10.0, "50"},
					{0.0, "60sec"},
				},
			},
			Series: []chart.Series{

				//RAM line
				chart.ContinuousSeries{
					Style: chart.Style{
						Show:        true,
						StrokeColor: drawing.ColorGreen,
						FillColor:   drawing.ColorGreen.WithAlpha(64),
					},
					XValues: xAxisPoints[:],
					YValues: ramPoints[:],
				},

				//CPU line
				chart.ContinuousSeries{
					Style: chart.Style{
						Show:        true,
						StrokeColor: drawing.ColorRed,
						FillColor:   drawing.ColorRed.WithAlpha(64),
					},
					XValues: xAxisPoints[:],
					YValues: cpuPoints[:],
				},

				//RAM annotation
				chart.AnnotationSeries{
					Annotations: []chart.Value2{
						{
							Style: chart.Style{
								StrokeColor: drawing.ColorGreen,
							},
							XValue: 60, YValue: m.UsedPercent, Label: "RAM " + FloatToString(m.UsedPercent) + "%"},
					},
				},

				//CPU annotation
				chart.AnnotationSeries{
					Annotations: []chart.Value2{
						{
							Style: chart.Style{
								StrokeColor: drawing.ColorRed,
							},
							XValue: 60, YValue: c[0], Label: "CPU " + FloatToString(c[0]) + "%"},
					},
				},
			},
		}

		imgBuffer := bytes.NewBuffer([]byte{})

		graph.Render(chart.PNG, imgBuffer)

		monChan <- imgBuffer

		ShiftPoints(cpuPoints[:])
		ShiftPoints(ramPoints[:])

		time.Sleep(DELAY)
	}
}
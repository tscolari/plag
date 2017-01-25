package output

import (
	"fmt"
	"time"

	"github.com/cloudflare/golibs/ewma"
	"github.com/gizak/termui"
	termbox "github.com/nsf/termbox-go"
	"github.com/tscolari/plag/parser"
)

type Graph struct {
	barChart  *termui.BarChart
	lineChart *termui.LineChart
	points    []parser.Data
	decay     time.Duration
}

func NewGraph(title string, decay time.Duration) *Graph {
	bc := termui.NewBarChart()
	bc.BorderLabel = title
	bc.TextColor = termui.ColorGreen
	bc.BarColor = termui.ColorBlue
	bc.NumColor = termui.ColorYellow
	termui.Render(bc)

	lc := termui.NewLineChart()
	lc.AxesColor = termui.ColorWhite
	lc.LineColor = termui.ColorYellow | termui.AttrBold
	lc.BorderLabel = fmt.Sprintf("EWMA %s", decay.String())
	termui.Render(lc)

	return &Graph{
		barChart:  bc,
		lineChart: lc,
		points:    []parser.Data{},
		decay:     decay,
	}
}

func (o *Graph) Write(data parser.Data) error {
	o.insertData(data)
	w, h := termbox.Size()
	o.barChart.Data, o.barChart.DataLabels = o.pointsToDataAndLabel(w)

	o.barChart.Width = w
	o.barChart.Height = h / 2
	o.barChart.X = 0
	o.barChart.Y = 0

	o.lineChart.Width = w
	o.lineChart.Height = h / 2
	o.lineChart.X = 0
	o.lineChart.Y = h / 2
	o.lineChart.Mode = "dot"
	o.lineChart.Data = o.pointsToEWMA(w)

	termui.Render(o.barChart, o.lineChart)
	return nil
}

func (o *Graph) Close() error {
	return nil
}

func (o *Graph) insertData(data parser.Data) {
	for i, p := range o.points {
		if p.Timestamp.UnixNano() > data.Timestamp.UnixNano() {
			o.points = append(o.points[:i], append([]parser.Data{data}, o.points[i:]...)...)
			return
		}
	}

	o.points = append(o.points, data)
}

func (o *Graph) pointsToDataAndLabel(maxWidth int) ([]int, []string) {
	data := []int{}
	labels := []string{}

	for i, p := range o.points {
		data = append(data, int(p.Value.Nanoseconds()))
		if len(labels) == 0 {
			labels = []string{"0"}
		} else {
			diff := fmt.Sprintf("%.1f", p.Timestamp.Sub(o.points[i-1].Timestamp).Seconds())
			labels = append(labels, diff)
		}
	}

	if len(data) > maxWidth/3 {
		return data[len(data)-maxWidth/3:], labels[len(labels)-maxWidth/3:]
	}

	return data, labels
}

func (o *Graph) pointsToEWMA(maxWidth int) []float64 {
	floats := []float64{}
	ewma := ewma.NewEwma(o.decay)

	for _, p := range o.points {
		floats = append(floats, ewma.Update(p.Value.Seconds(), p.Timestamp))
	}

	if len(floats) > maxWidth {
		return floats[len(floats)-maxWidth:]
	}

	return floats
}

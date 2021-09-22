package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/jonfriesen/finance-go"
	"github.com/pkg/errors"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func Sparkline(symbol string, charBars []*finance.ChartBar) (*bytes.Buffer, error) {
	var dates []time.Time
	var yv []float64
	for _, ts := range charBars {
		parsed := time.Unix(int64(ts.Timestamp), 0)
		dates = append(dates, parsed)

		close, _ := ts.Close.Float64()
		yv = append(yv, close)
	}

	priceSeries := chart.TimeSeries{
		XValues: dates,
		YValues: yv,
	}

	graph := chart.Chart{
		Title: fmt.Sprintf("%s - 30d", symbol),
		TitleStyle: chart.Style{
			Show:      true,
			FontColor: drawing.ColorFromHex("374151").WithAlpha(80),
		},
		Width:  300,
		Height: 45,
		XAxis: chart.XAxis{
			ValueFormatter: chart.TimeHourValueFormatter,
		},
		Series: []chart.Series{
			priceSeries,
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.SVG, buffer)
	if err != nil {
		return nil, errors.Wrap(err, "rendering chart image")
	}

	return buffer, nil
}

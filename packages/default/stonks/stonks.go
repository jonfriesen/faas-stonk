package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jonfriesen/finance-go"
	"github.com/jonfriesen/finance-go/chart"
	"github.com/jonfriesen/finance-go/datetime"
	"github.com/jonfriesen/finance-go/quote"
	"github.com/pkg/errors"
)

type MarketState string

type YQuote struct {
	Chart           string
	Trend           string // can be bigup (>3%), up, drop or bigdrop (<3%)
	Symbol          string
	Name            string
	Currency        string
	MarketState     string
	MarketPrice     float64
	MarketChange    float64
	MarketChangePct float64
}

// This main function is for testing locally
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		qp := r.URL.Query()

		params := make(map[string]interface{}, len(qp))
		for k, v := range qp {
			params[k] = v
		}
		w.Header().Set("Content-Type", "text/html")
		payload := Main(params)
		if b, ok := payload["body"]; ok {
			fmt.Fprint(w, b)
		}
		if b, ok := payload["error"]; ok {
			fmt.Fprint(w, b)
		}
	})

	// push up to web?
	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Running on port :8080")
	http.ListenAndServe(":8080", nil)
}

func Main(args map[string]interface{}) map[string]interface{} {
	symbol, ok := args["symbol"].(string)
	if !ok {
		symbol = "DOCN"
	}

	resp := make(map[string]interface{})

	q := getQuote(symbol)

	// b, err := json.Marshal(q)
	// if err != nil {
	// 	resp["error"] = errors.Wrap(err, "marshalling quote object")
	// 	return resp
	// }

	start := time.Now().Add(time.Hour * 24 * 30 * -1)
	end := time.Now()

	history := []*finance.ChartBar{}
	iter := chart.Get(&chart.Params{
		Symbol:   symbol,
		Interval: datetime.OneDay,
		Start:    datetime.New(&start),
		End:      datetime.New(&end),
	})

	for iter.Next() {
		b := iter.Bar()
		history = append(history, b)
	}
	if err := iter.Err(); err != nil {
		resp["error"] = errors.Wrap(err, "getting history")
		return resp
	}

	sb, err := Sparkline(symbol, history)
	if err != nil {
		resp["error"] = errors.Wrap(err, "getting sparkline")
		return resp
	}

	q.Chart = sb.String()

	b, err := GetHTML(q)
	if err != nil {
		resp["error"] = errors.Wrap(err, "getting html")
		return resp
	}

	resp["body"] = b
	return resp
}

func getQuote(symbol string) YQuote {
	var yq YQuote

	var MarketPrice float64
	var MarketChange float64
	var MarketChangePct float64

	q, err := quote.Get(symbol)
	if q == nil || err != nil {
		yq = YQuote{
			Symbol:      symbol,
			Trend:       "?",
			MarketState: "?",
		}
	} else {
		if q.MarketState == "PRE" {
			MarketPrice = q.PreMarketPrice
			MarketChange = q.PreMarketChange
			MarketChangePct = q.PreMarketChangePercent

		} else if q.MarketState == "POST" {
			MarketPrice = q.PostMarketPrice
			MarketChange = q.PostMarketChange
			MarketChangePct = q.PostMarketChangePercent
		} else {
			MarketPrice = q.RegularMarketPrice
			MarketChange = q.RegularMarketChange
			MarketChangePct = q.RegularMarketChangePercent
		}
		yq = YQuote{
			Symbol:          q.Symbol,
			Name:            q.ShortName,
			Currency:        q.CurrencyID,
			Trend:           GetTrendIcon(GetTrend(MarketChangePct)),
			MarketState:     string(q.MarketState),
			MarketPrice:     MarketPrice,
			MarketChange:    MarketChange,
			MarketChangePct: MarketChangePct,
		}

	}
	return yq
}

func GetMarketIcon(state string) string {
	states := map[string]string{
		"PRE":     "⏭",
		"REGULAR": "▶",
		"POST":    "⏮",
		"?":       "?",
	}
	if icon, ok := states[state]; ok {
		return icon
	} else {
		return "⏹"
	}
}

func GetTrendIcon(trend string) string {
	icons := map[string]string{
		"bigup":   "⬆️ ",
		"up":      "↗️ ",
		"drop":    "↘️ ",
		"bigdrop": "⬇️ ",
	}
	return icons[trend]
}

func GetTrend(pct float64) string {
	var trend string
	if pct > 3 {
		trend = "bigup"
	} else if pct > 0 {
		trend = "up"
	} else if pct > -3 {
		trend = "drop"
	} else {
		trend = "bigdrop"
	}
	return trend
}

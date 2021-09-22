package main

import (
	"bytes"
	"html/template"
	"strconv"
)

var fmap template.FuncMap

func init() {
	fmap = template.FuncMap{}
	fmap["comma"] = func(n int) string {
		in := strconv.FormatInt(int64(n), 10)
		numOfDigits := len(in)
		if n < 0 {
			numOfDigits-- // First character is the - sign (not a digit)
		}
		numOfCommas := (numOfDigits - 1) / 3

		out := make([]byte, len(in)+numOfCommas)
		if n < 0 {
			in, out[0] = in[1:], '-'
		}

		for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
			out[j] = in[i]
			if i == 0 {
				return string(out)
			}
			if k++; k == 3 {
				j, k = j-1, 0
				out[j] = ','
			}
		}
	}
	fmap["safeHTML"] = func(s string) template.HTML { return template.HTML(s) }
}

const quoteTmpl = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
</head>
<body>
<h1><b>{{.Symbol}}</b> {{.Name}}</h1>
{{safeHTML .Chart}}
<ul>
    <li>${{printf "%.2f" .MarketPrice}}<small>{{.Currency}}</small> {{.Trend}}</li>
	<li>Change ${{printf "%.2f" .MarketChange}} ({{printf "%.2f" .MarketChangePct}}%)</li>
</ul>
</body>
</html>`

func GetHTML(quote YQuote) (string, error) {
	t, err := template.New("symbol").Funcs(fmap).Parse(quoteTmpl)
	if err != nil {
		return "", err
	}

	var f bytes.Buffer
	err = t.Execute(&f, quote)
	if err != nil {
		return "", err
	}

	return f.String(), nil
}

package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// Parse parses args from request to make list of strings (e.g. BTC-USDT -> BTCUSDT)
func ParseString(s string) []string {
	f := strings.Split(s, ",")
	newStr := []string{}

	for _, val := range f {
		newStr = append(newStr, strings.Replace(val, "-", "", 1))
	}

	return newStr
}

// MakeString making string valid for sending to Binance (e.g. ["BTCUSDT","ETHUSDT"])
func MakeString(s []string) string {
	str := `[`

	for i, val := range s {
		str += `"` + val + `"`
		if i != len(s)-1 {
			str += `,`
		}
	}

	str += `]`

	return str
}

func ParseRequest(body io.ReadCloser) (string, error) {
	resBody, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}

type Rate struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func ParseRate(resp *http.Response) (Rate, error) {
	var r Rate

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Rate{}, err
	}

	//fmt.Println(string(body))
	//var s string
	if err := json.Unmarshal(body, &r); err != nil {
		return Rate{}, err
	}

	return r, nil
}

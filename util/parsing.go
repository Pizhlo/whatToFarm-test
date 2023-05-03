package util

import (
	"encoding/json"
	"fmt"
	"io"
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

// Rate saves a response from server
type Rate struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// ParseRate parses response into Rate structure
func ParseRate(body io.ReadCloser) ([]Rate, error) {
	var r []Rate

	res, err := io.ReadAll(body)
	if err != nil {
		fmt.Println("unable to read: ", err)
		return []Rate{}, err
	}

	if err := json.Unmarshal(res, &r); err != nil {
		return []Rate{}, err
	}

	return r, nil
}

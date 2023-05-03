package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

func ParseRequest(body io.ReadCloser) string {
	resBody, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	return string(resBody)
}

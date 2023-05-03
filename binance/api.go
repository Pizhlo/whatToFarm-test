package binance

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 OPR/97.0.0.0"

// MakeRequest makes a request using url, method and args and returns body of response
func MakeRequest(method string, url string, args string) io.ReadCloser {
	req, err := http.NewRequest(method, fmt.Sprintf(url, args), nil)
	if err != nil {
		fmt.Fprint(os.Stderr, "Unable to make request:", err)
		os.Exit(1)
	}

	req.Header = http.Header{
		"User-Agent": {UserAgent},
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprint(os.Stderr, "Unable to call Binance API:", err)
		os.Exit(1)
	}

	return resp.Body
}

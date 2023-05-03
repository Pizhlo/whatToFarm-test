package client

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Pizhlo/whatToFarm-test/util"
)

// MakeRequest makes a request to server with the specified arguments 
func MakeRequest(pairs string) {

	url := fmt.Sprintf("http://localhost:3001/api/v1/rates?pairs=%s", pairs)

	r, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, "Unable to make request:", err)
		os.Exit(1)
	}

	resp, err := util.ParseRate(r.Body)
	if err != nil {
		fmt.Printf("could not read response body: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(resp[0].Price)

}

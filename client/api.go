package client

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Pizhlo/whatToFarm-test/server"
	"github.com/Pizhlo/whatToFarm-test/util"
)

func MakeRequest(pairs string) {

	url := fmt.Sprintf("http://localhost:3001/api/v1/rates?pairs=%s", pairs)
	method := server.MethodGet

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Fprint(os.Stderr, "Unable to make request:", err)
		os.Exit(1)
	}

	client := &http.Client{}
	body, err := client.Do(req)
	if err != nil {
		fmt.Fprint(os.Stderr, "Unable to call server:", err)
		os.Exit(1)
	}

	resp, err := util.ParseRate(body)
	if err != nil {
		fmt.Printf("could not read response body: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(resp)

}

package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/Pizhlo/whatToFarm-test/binance"
	"github.com/Pizhlo/whatToFarm-test/util"
)

const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

type RatesRequest struct {
	Pairs string `form:"pairs"`
}

// GetRates parses args from request and makes request to Binance API
func GetRates(ctx *gin.Context) {
	var req RatesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	parsedString := util.ParseString(req.Pairs)
	args := util.MakeString(parsedString)
	url := `https://api.binance.com/api/v3/ticker/price?symbols=%s`
	body := binance.MakeRequest(MethodGet, url, args)
	resp, err := util.ParseRequest(body)
	if err != nil {
		fmt.Printf("could not read response body: %s\n", err)
		os.Exit(1)
	}
	ctx.JSON(http.StatusOK, resp)
}

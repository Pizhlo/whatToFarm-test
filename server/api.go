package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Pizhlo/whatToFarm-test/binance"
	"github.com/Pizhlo/whatToFarm-test/util"
)

const (
	methodGet  = "GET"
	methodPost = "POST"
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
	body := binance.MakeRequest(methodGet, url, args)
	resp := util.ParseRequest(body)
	ctx.JSON(http.StatusOK, resp)
}

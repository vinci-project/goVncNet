package tools

import (
	"encoding/json"
	"goVncNet/helpers"
	"strconv"

	"github.com/valyala/fasthttp"
)

func MakeResponse(statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
}

func MakeBalanceResponse(balances map[string]string,
	stakes map[string]string,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	response := helpers.BalanceResponse{balances, stakes}
	jsResponse, _ := json.Marshal(response)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(jsResponse)
}

func MakeStakeResponse(stakes map[string]string,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	response := helpers.StakeResponse{stakes}
	jsResponse, _ := json.Marshal(response)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(jsResponse)
}

func MakeDataResponse(data string,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody([]byte(data))
}

func MakeVersionResponse(version string,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	response := helpers.VersionResponse{version}
	jsResponse, _ := json.Marshal(response)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(jsResponse)
}

func MakeBHeightResponse(bheight string,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	response := helpers.BHeightResponse{bheight}
	jsResponse, _ := json.Marshal(response)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(jsResponse)
}

func MakeASResponse(applicants []string,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	response := helpers.ASResponse{applicants}
	jsResponse, _ := json.Marshal(response)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(jsResponse)
}

func MakeVSResponse(votes string,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	response := helpers.VSResponse{votes}
	jsResponse, _ := json.Marshal(response)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(jsResponse)
}

func MakeAVSResponse(votes map[string]string,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	response := helpers.AVSResponse{votes}
	jsResponse, _ := json.Marshal(response)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(jsResponse)
}

func MakeStatisticsResponse(status bool,
	tranCount int64,
	ltps int64,
	blockHeight int64,
	ltpb int64,
	tpd float64,
	upd int64,
	statusCode int,
	ctx *fasthttp.RequestCtx) {
	//

	response := helpers.StatisticsResponse{strconv.FormatInt(tranCount, 10),
		strconv.FormatInt(ltps, 10),
		strconv.FormatInt(blockHeight, 10),
		strconv.FormatInt(ltpb, 10),
		strconv.FormatFloat(tpd, 'f', 8, 64),
		[]string{"50.11", "8.68"},
		strconv.FormatInt(upd, 10)}

	jsResponse, _ := json.Marshal(response)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(jsResponse)
}

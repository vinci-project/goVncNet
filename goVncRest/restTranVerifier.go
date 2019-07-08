package restServer

import (
	"goVncNet/helpers"
	"log"
	"strconv"

	"github.com/valyala/fasthttp"
)

func verifyPostTransaction(ctx *fasthttp.RequestCtx) (statusCode int,
	transactionForDB []byte,
	transactionTime int64) {
	//

	transactionForDB = ctx.PostBody()
	transactionForDBString := string(transactionForDB)
	log.Println(transactionForDBString)
	tranType := helpers.GetRawTransactionType(transactionForDBString)
	switch tranType {
	//

	case "ST":
		simpleTran, err := helpers.ParseSimpleTransaction(transactionForDBString)
		if err != nil {
			//

			return helpers.StatusWrongDataFormat, transactionForDB, transactionTime
		}

		transactionTime, statusCode, ok := helpers.VerifySimpleTransaction(simpleTran)
		if !ok {
			//

			return statusCode, transactionForDB, transactionTime
		}

		return helpers.StatusOk, transactionForDB, transactionTime

	case "AT":
		simpleTran, err := helpers.ParseApplicantTransaction(transactionForDBString)
		if err != nil {
			//

			return helpers.StatusWrongDataFormat, transactionForDB, transactionTime
		}

		transactionTime, statusCode, ok := helpers.VerifyApplicantTransaction(simpleTran)
		if !ok {
			//

			return statusCode, transactionForDB, transactionTime
		}

		return helpers.StatusOk, transactionForDB, transactionTime

	case "VT":
		simpleTran, err := helpers.ParseVoteTransaction(transactionForDBString)
		if err != nil {
			//

			return helpers.StatusWrongDataFormat, transactionForDB, transactionTime
		}

		transactionTime, statusCode, ok := helpers.VerifyVoteTransaction(simpleTran)
		if !ok {
			//

			return statusCode, transactionForDB, transactionTime
		}

		return helpers.StatusOk, transactionForDB, transactionTime

	case "UAT":
		simpleTran, err := helpers.ParseUATransaction(transactionForDBString)
		if err != nil {
			//

			return helpers.StatusWrongDataFormat, transactionForDB, transactionTime
		}

		transactionTime, statusCode, ok := helpers.VerifyUATransaction(simpleTran)
		if !ok {
			//

			return statusCode, transactionForDB, transactionTime
		}

		return helpers.StatusOk, transactionForDB, transactionTime

	case "UVT":
		simpleTran, err := helpers.ParseUVTransaction(transactionForDBString)
		if err != nil {
			//

			return helpers.StatusWrongDataFormat, transactionForDB, transactionTime
		}

		transactionTime, statusCode, ok := helpers.VerifyUVTransaction(simpleTran)
		if !ok {
			//

			return statusCode, transactionForDB, transactionTime
		}

		return helpers.StatusOk, transactionForDB, transactionTime

	}

	return helpers.StatusUnknownTranType, transactionForDB, transactionTime
}

func verifyBalanceRequest(ctx *fasthttp.RequestCtx) (statusCode int, senders []string, ttoken string) {
	//

	args := ctx.QueryArgs()
	for errNum, v := range helpers.RequestBalanceFields {
		//

		if !args.Has(v) {
			//

			return errNum, senders, ttoken
		}
	}

	ttoken = string(args.Peek("TTOKEN"))
	sendersBytes := ctx.QueryArgs().PeekMulti("SENDER")
	for i := range sendersBytes {
		//

		sender := string(sendersBytes[i])
		if len(sender) != 66 {
			//

			return helpers.StatusWrongAttr_SENDER, senders, ttoken
		}

		senders = append(senders, sender)
	}

	return helpers.StatusOk, senders, ttoken
}

func verifyStakeRequest(ctx *fasthttp.RequestCtx) (statusCode int, senders []string) {
	//

	args := ctx.QueryArgs()
	for errNum, v := range helpers.RequestStakeFields {
		//

		if !args.Has(v) {
			//

			return errNum, senders
		}
	}

	sendersBytes := ctx.QueryArgs().PeekMulti("SENDER")
	for i := range sendersBytes {
		//

		sender := string(sendersBytes[i])
		if len(sender) != 66 {
			//

			return helpers.StatusWrongAttr_SENDER, senders
		}

		senders = append(senders, sender)
	}

	return helpers.StatusOk, senders
}

func verifyTranStatusRequest(ctx *fasthttp.RequestCtx) (statusCode int, key string) {
	//

	args := ctx.QueryArgs()
	for errNum, v := range helpers.RequestTranStatusFields {
		//

		if !args.Has(v) {
			//

			return errNum, key
		}
	}

	key = string(args.Peek("KEY"))

	if len(key) != 64 {
		//

		return helpers.StatusWrongAttr_KEY, key
	}

	return helpers.StatusOk, key
}

func verifyVSRequest(ctx *fasthttp.RequestCtx) (statusCode int, address string) {
	//

	args := ctx.QueryArgs()
	for errNum, v := range helpers.RequestVSFields {
		//

		if !args.Has(v) {
			//

			return errNum, address
		}
	}

	address = string(args.Peek("ADDRESS"))

	if len(address) != 66 {
		//

		return helpers.StatusWrongAttr_ADDRESS, address
	}

	return helpers.StatusOk, address
}

func verifyGetBlockRequest(ctx *fasthttp.RequestCtx) (statusCode int, bheight int64) {
	//

	args := ctx.QueryArgs()
	for errNum, v := range helpers.RequestGetBlockFields {
		//

		if !args.Has(v) {
			//

			return errNum, bheight
		}
	}

	heightString := string(args.Peek("BHEIGHT"))
	bheight, err := strconv.ParseInt(heightString, 10, 64)
	if err != nil {
		//

		return helpers.StatusWrongAttr_BHEIGHT, bheight
	}

	return helpers.StatusOk, bheight
}

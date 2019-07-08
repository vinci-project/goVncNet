package restServer

import (
	"goVncNet/goVncRest/tools"
	"goVncNet/helpers"
	"net"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/valyala/fasthttp"
)

var redisDB *redis.Client
var tranChannel *chan string

func fastHTTPRawHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetConnectionClose()
	if string(ctx.Method()) == "GET" {
		//

		switch string(ctx.Path()) {

		case "/wallet/getBalance":
			statusCode, senders, ttoken := verifyBalanceRequest(ctx)
			if statusCode >= 600 {
				//

				tools.MakeResponse(statusCode, ctx)
				return

			} else {
				//

				balances := make(map[string]string)
				stakes := make(map[string]string)
				for i := range senders {
					//

					var aStake int64 = 0
					var vStake int64 = 0
					var allStake int64 = 0
					zScore := redisDB.ZScore("APPLICANTS", senders[i])
					if !helpers.IsRedisError(zScore) {
						//

						aStake = 16000
					}

					zScore = redisDB.ZScore("UNTVOTES", senders[i])
					if !helpers.IsRedisError(zScore) {
						//

						vStake = int64(zScore.Val()) * 10
					}

					allStake = aStake + vStake

					zScore = redisDB.ZScore("BALANCE:"+ttoken, senders[i])
					if helpers.IsRedisError(zScore) {
						//

						balances[senders[i]] = "0"

					} else {
						//

						balances[senders[i]] = strconv.FormatFloat(zScore.Val()-float64(allStake), 'f', -1, 64)
					}

					stakes[senders[i]] = strconv.FormatInt(allStake, 10)
				}

				tools.MakeBalanceResponse(balances, stakes, helpers.StatusOk, ctx)
				return
			}

		case "/wallet/getStake":
			statusCode, senders := verifyStakeRequest(ctx)
			if statusCode >= 600 {
				//

				tools.MakeResponse(statusCode, ctx)
				return

			} else {
				//

				stakes := make(map[string]string)
				for i := range senders {
					//

					var aStake int64 = 0
					var vStake int64 = 0
					zScore := redisDB.ZScore("APPLICANTS", senders[i])
					if !helpers.IsRedisError(zScore) {
						//

						aStake = 16000
					}

					zScore = redisDB.ZScore("UNTVOTES", senders[i])
					if !helpers.IsRedisError(zScore) {
						//

						vStake = int64(zScore.Val()) * 10
					}

					stakes[senders[i]] = strconv.FormatInt(aStake+vStake, 10)
				}

				tools.MakeStakeResponse(stakes, helpers.StatusOk, ctx)
				return
			}

		case "/wallet/tranStatus":
			statusCode, key := verifyTranStatusRequest(ctx)
			if statusCode >= 600 {
				//

				tools.MakeResponse(statusCode, ctx)
				return

			} else {
				//

				zScore := redisDB.ZScore("COMPLETE TRANSACTIONS", key)
				if !helpers.IsRedisError(zScore) {
					//

					tools.MakeResponse(helpers.StatusOk, ctx)
					return
				}

				zScore = redisDB.ZScore("FAILED TRANSACTIONS", key)
				if !helpers.IsRedisError(zScore) {
					//

					tools.MakeResponse(helpers.StatusTranFailed, ctx)
					return
				}

				tools.MakeResponse(helpers.StatusTranNotFound, ctx)
				return
			}

		case "/blockchain/getBHeight":
			intCmd := redisDB.ZCard("VNCCHAIN")
			if helpers.IsRedisError(intCmd) {
				//

				tools.MakeResponse(helpers.StatusInternalServerError, ctx)
				return
			}

			tools.MakeBHeightResponse(strconv.FormatInt(intCmd.Val(), 10), helpers.StatusOk, ctx)

		case "/blockchain/getTran":
			statusCode, key := verifyTranStatusRequest(ctx)
			if statusCode >= 600 {
				//

				tools.MakeResponse(statusCode, ctx)
				return

			} else {
				//

				stringCmd := redisDB.Get("TRANSACTIONS:" + key)
				if helpers.IsRedisError(stringCmd) {
					//

					tools.MakeResponse(helpers.StatusTranNotFound, ctx)
					return
				}

				tools.MakeDataResponse(stringCmd.Val(), helpers.StatusOk, ctx)
				return
			}

		case "/blockchain/getBlock":
			statusCode, bheight := verifyGetBlockRequest(ctx)
			if statusCode >= 600 {
				//

				tools.MakeResponse(statusCode, ctx)
				return

			} else {
				//

				stringSliceCmd := redisDB.ZRange("VNCCHAIN", bheight-1, bheight)
				if helpers.IsRedisError(stringSliceCmd) {
					//

					tools.MakeResponse(helpers.StatusDataNotFound, ctx)
					return
				}

				if len(stringSliceCmd.Val()) != 1 {
					//

					tools.MakeResponse(helpers.StatusDataNotFound, ctx)
					return
				}

				tools.MakeDataResponse(stringSliceCmd.Val()[0], helpers.StatusOk, ctx)
				return
			}

		case "/blockchain/getVersion":
			stringCmd := redisDB.Get("VERSION")
			if helpers.IsRedisError(stringCmd) {
				//

				tools.MakeResponse(helpers.StatusDataNotFound, ctx)
				return
			}

			tools.MakeVersionResponse(stringCmd.Val(), helpers.StatusOk, ctx)
			return

		case "/blockchain/getNodes":
			stringCmd := redisDB.Get("NODES LIST")
			if helpers.IsRedisError(stringCmd) {
				//

				tools.MakeResponse(helpers.StatusDataNotFound, ctx)
				return
			}

			tools.MakeDataResponse(stringCmd.Val(), helpers.StatusOk, ctx)
			return

		case "/blockchain/as":
			stringSliceCmd := redisDB.ZRange("APPLICANTS", 0, -1)
			if helpers.IsRedisError(stringSliceCmd) {
				//

				tools.MakeResponse(helpers.StatusDataNotFound, ctx)
				return
			}

			tools.MakeASResponse(stringSliceCmd.Val(), helpers.StatusOk, ctx)
			return

		case "/blockchain/vs":
			statusCode, address := verifyVSRequest(ctx)
			if statusCode >= 600 {
				//

				tools.MakeResponse(statusCode, ctx)
				return

			} else {
				//

				zScore := redisDB.ZScore("UNTVOTES", address)
				if helpers.IsRedisError(zScore) {
					//

					tools.MakeResponse(helpers.StatusDataNotFound, ctx)
					return
				}

				tools.MakeVSResponse(strconv.FormatFloat(zScore.Val(), 'f', -1, 64), helpers.StatusOk, ctx)
				return
			}

		case "/blockchain/avs":
			stringSliceCmd := redisDB.ZRange("APPLICANTS", 0, -1)
			if helpers.IsRedisError(stringSliceCmd) {
				//

				tools.MakeResponse(helpers.StatusDataNotFound, ctx)
				return
			}

			addresses := stringSliceCmd.Val()
			votes := make(map[string]float64)
			for i := range addresses {
				//

				votes[addresses[i]] = 0
				stringSliceCmd = redisDB.ZRange("VOTES:"+addresses[i], 0, -1)
				if len(stringSliceCmd.Val()) == 0 {
					//

					continue
				}

				for _, j := range stringSliceCmd.Val() {
					//

					zScore := redisDB.ZScore("VOTES:"+addresses[i], j)
					if helpers.IsRedisError(zScore) {
						//

						continue
					}

					votes[addresses[i]] += zScore.Val()
				}
			}

			votesString := make(map[string]string)
			for k, v := range votes {
				//

				votesString[k] = strconv.FormatFloat(v, 'f', -1, 64)
			}

			tools.MakeAVSResponse(votesString, helpers.StatusOk, ctx)
			return

		default:
			//

			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}

		return

	} else if string(ctx.Method()) == "POST" {
		//

		switch string(ctx.Path()) {

		case "/wallet/transaction":
			statusCode, transactionForDb, transactionTime := verifyPostTransaction(ctx)
			if statusCode >= 600 {
				//

				tools.MakeResponse(statusCode, ctx)
				return

			} else {
				//

				errRedis := redisDB.ZAdd("RAW TRANSACTIONS", redis.Z{
					Score:  float64(transactionTime),
					Member: transactionForDb,
				})

				if helpers.IsRedisError(errRedis) {
					//

					tools.MakeResponse(helpers.StatusInternalServerError, ctx)
					return
				}

				*tranChannel <- string(transactionForDb)

				tools.MakeResponse(helpers.StatusOk, ctx)
				return
			}
		}
	}

	ctx.Error("Unsupported method", fasthttp.StatusMethodNotAllowed)
}

func Start(r *redis.Client, c *chan string, ip string) {
	//

	redisDB = r
	tranChannel = c

	// listener, err := reuseport.Listen("tcp4", net.JoinHostPort(ip, "5000"))
	// if err != nil {
	// 	log.Fatalf("error in reuseport listener: %s", err)
	// }

	server := &fasthttp.Server{
		Handler:          fastHTTPRawHandler,
		DisableKeepalive: true,
	}

	panic(server.ListenAndServe(net.JoinHostPort(ip, "5000")))
}

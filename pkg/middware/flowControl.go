package middware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/KenLu0929/flowControlTask/config/rdb"
	"github.com/KenLu0929/flowControlTask/pkg/models"

	"github.com/gin-gonic/gin"
)

var redisDB = rdb.Rdb
var ctx = rdb.Ctx

const resetTime = 3600 //1 hour
const max = 1000       //request max time

func FlowControlByIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIP := c.ClientIP()
		now := time.Now().Unix()

		var count int64
		var reset int64

		val, err := redisDB.HMGet(ctx, userIP, "count", "reset").Result()
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.Result{
				Success: false,
				Error:   err.Error(),
			})
			c.Abort()
			return
		}

		//key no exists
		if val[0] == nil || val[1] == nil {
			count = 1
			reset = now + resetTime
			redisDB.HMSet(ctx, userIP, "count", count, "reset", reset)
		} else {
			val_count, _ := strconv.ParseInt(val[0].(string), 10, 64)
			val_reset, _ := strconv.ParseInt(val[1].(string), 10, 64)

			if val_reset < now {
				count = 1
				reset = now + resetTime
				redisDB.HMSet(ctx, userIP, "count", count, "reset", reset)
			} else if val_count == max {
				c.JSON(http.StatusTooManyRequests, models.Result{
					Success: false,
					Error:   "too many request",
				})
				c.Abort()
				return
			} else {
				reset = val_reset
				count = val_count + 1
				redisDB.HIncrBy(ctx, userIP, "count", 1)
			}
		}

		c.Header("X-RateLimit-Remaining", strconv.FormatInt(max-count, 10))
		c.Header("X-RateLimit-Reset", time.Unix(reset, 0).Format("2006-01-02 15:04:05"))
		c.Next()
	}
}

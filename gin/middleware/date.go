package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/zerosnake0/goutil/zerolog"
)

func DateValidatorMiddleware(maxDur time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		date := c.GetHeader("Date")
		if date == "" {
			zerolog.FromGinCtx(c).Error().Msg("empty date")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		t, err := time.ParseInLocation(time.RFC1123, date, time.UTC)
		if err != nil {
			zerolog.FromGinCtx(c).Error().Str("date", date).Msg("bad date")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		diff := time.Now().Sub(t)
		if diff < 0 {
			diff = -diff
		}
		if diff > maxDur {
			zerolog.FromGinCtx(c).Error().Dur("diff", diff).
				Dur("max", maxDur).
				Msg("date difference exceeds limit")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Next()
	}
}

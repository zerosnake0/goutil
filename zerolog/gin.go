package zerolog

import (
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const zerologKey = "_goutil_zerolog"

var disabled *zerolog.Logger

func init() {
	nop := zerolog.Nop()
	disabled = &nop
}

func FromGinCtx(c *gin.Context) *zerolog.Logger {
	o, ok := c.Get(zerologKey)
	if !ok {
		return disabled
	}
	l, ok := o.(*zerolog.Logger)
	if !ok {
		return disabled
	}
	return l
}

func setLogger(c *gin.Context, logger *zerolog.Logger) {
	c.Set(zerologKey, logger)
}

func updateLoggerStr(c *gin.Context, key, val string) {
	l := FromGinCtx(c)
	l.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str(key, val)
	})
}

func NewHandler(logger *zerolog.Logger) gin.HandlerFunc {
	if logger == nil {
		logger = &log.Logger
	}
	return func(c *gin.Context) {
		l := logger.With().Logger()
		setLogger(c, &l)
		c.Next()
	}
}

func URLHandler(fieldKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		updateLoggerStr(c, fieldKey, c.Request.URL.String())
		c.Next()
	}
}

func MethodHandler(fieldKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		updateLoggerStr(c, fieldKey, c.Request.Method)
		c.Next()
	}
}

func RequestHandler(fieldKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		updateLoggerStr(c, fieldKey, c.Request.Method+" "+c.Request.URL.String())
		c.Next()
	}
}

func RemoteAddrHandler(fieldKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if host, _, err := net.SplitHostPort(c.Request.RemoteAddr); err == nil {
			updateLoggerStr(c, fieldKey, host)
		}
		c.Next()
	}
}

func UserAgentHandler(fieldKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if ua := c.GetHeader("User-Agent"); ua != "" {
			updateLoggerStr(c, fieldKey, ua)
		}
		c.Next()
	}
}

func RefererHandler(fieldKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if ref := c.GetHeader("Referer"); ref != "" {
			updateLoggerStr(c, fieldKey, ref)
		}
		c.Next()
	}
}

func RequestIDHandler(fieldKey, headerName string) gin.HandlerFunc {
	if headerName == "" {
		headerName = "X-Request-Id"
	}
	return func(c *gin.Context) {
		reqID := c.GetHeader(headerName)
		if reqID == "" {
			reqID = xid.New().String()
		}
		updateLoggerStr(c, fieldKey, reqID)
		c.Writer.Header().Set(headerName, reqID)
		c.Next()
	}
}

func AccessHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		FromGinCtx(c).Info().Dur("latency", latency).
			Str("client_ip", c.ClientIP()).
			Int("status_code", c.Writer.Status()).
			Str("error_msg", c.Errors.ByType(gin.ErrorTypePrivate).String()).
			Int("body_size", c.Writer.Size()).
			Msg("")
	}
}

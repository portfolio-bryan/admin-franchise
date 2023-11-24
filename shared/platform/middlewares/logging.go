package middlewares

import (
	"github.com/bperezgo/admin_franchise/shared/platform/handlertypes"
	"github.com/bperezgo/admin_franchise/shared/platform/logger"
	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestID string

		ctx := c.Request.Context()
		ctx.Value(RequestIDKey)

		requestID, ok := ctx.Value(RequestIDKey).(string)
		if !ok {
			requestID = ""
		}

		l := logger.GetLogger()
		l.Info(logger.LogInput{
			Action: "REQUEST",
			State:  logger.SUCCESS,
			Http: &logger.LogHttpInput{
				Request: handlertypes.Request{},
			},
			Meta: &handlertypes.Meta{
				RequestId: requestID,
			},
		})

		c.Next()

		l.Info(logger.LogInput{
			Action: "RESPONSE",
			State:  logger.SUCCESS,
			Http:   &logger.LogHttpInput{},
			Meta: &handlertypes.Meta{
				RequestId: requestID,
			},
		})

	}
}

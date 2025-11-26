package middleware

import (
	"net"
	"net/http"
	"os"
	"strings"
	"xbase/errors"
	"xbase/log"
	"xrobot/internal/http/response"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if err == response.ExitPanic {
					c.AbortWithStatus(http.StatusOK)
					return
				}

				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				log.Panic(err)

				if brokenPipe {
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) //nolint: errcheck
					c.Abort()
				} else {
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()

		c.Next()
	}
}

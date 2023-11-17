package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"

	"github.com/nicelizhi/easy-admin-core/sdk/config"
)

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.SslConfig.Enable {
			secureMiddleware := secure.New(secure.Options{
				SSLRedirect:   true,
				SSLHost:       config.SslConfig.Domain,
				IsDevelopment: true,
			})
			err := secureMiddleware.Process(c.Writer, c.Request)
			if err != nil {
				c.Abort()
				return
			}
		} else {
			secureMiddleware := secure.New(secure.Options{
				IsDevelopment: true,
			})
			err := secureMiddleware.Process(c.Writer, c.Request)
			if err != nil {
				c.Abort()
				return
			}
		}

		// Avoid header rewrite if response is a redirection.
		if status := c.Writer.Status(); status > 300 && status < 399 {
			c.Abort()
		}

		c.Next()
	}
}

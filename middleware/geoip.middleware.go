package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/mr687/privy-be-test-go/service"
)

func GeoIpMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		realIp := service.GetRealIp(c)
		geoip, err := service.GetGeoInfo(realIp)
		if err != nil {
			c.Set("geoip", &service.GeoIp{})
		}

		fmt.Println(geoip)
		c.Set("geoip", geoip)
		c.Next()
	}
}

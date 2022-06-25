package membership

import (
	"github.com/labstack/echo/v4"
	"log"
)

var App = &Application{}

func InitMembershipRouter(e *echo.Echo) *echo.Group {
	App = NewApplication(*NewRepository(map[string]Membership{}))
	membershipGroup := e.Group("/memberships")
	membershipGroup.Use(LoggingMiddleware)
	membershipGroup.POST("", Create)
	membershipGroup.PATCH("", Update)
	membershipGroup.DELETE("/:id", Delete)
	membershipGroup.GET("/:id", Read)
	membershipGroup.GET("", ReadAll)
	return membershipGroup
}

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("====request====")
		log.Println("url:", c.Request().URL.String())
		log.Println("method:", c.Request().Method)
		log.Println("body:", c.Request().Body)
		log.Println("====response====")
		log.Println("status:", c.Response().Status)
		return next(c)
	}
}

package membership

import (
	"github.com/labstack/echo/v4"
	"log"
)

var App = &Application{}

func SetMembershipRouter(e *echo.Echo, controller Controller) *echo.Echo {
	App = NewApplication(*NewRepository(map[string]Membership{}))
	membershipGroup := e.Group("/memberships")
	membershipGroup.Use(LoggingMiddleware)
	membershipGroup.POST("", controller.Create)
	membershipGroup.PATCH("", controller.Update)
	membershipGroup.DELETE("/:id", controller.Delete)
	membershipGroup.GET("/:id", controller.Read)
	membershipGroup.GET("", controller.ReadAll)
	return e
}

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("====request====")
		log.Println("url:", c.Request().URL.String())
		log.Println("method:", c.Request().Method)
		log.Println("body:", c.Request().Body)
		err := next(c)
		log.Println("====response====")
		log.Println("status:", c.Response().Status)
		return err
	}
}

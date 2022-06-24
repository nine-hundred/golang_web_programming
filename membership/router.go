package membership

import (
	"github.com/labstack/echo/v4"
)

var App = &Application{}

func InitMembershipRouter(e *echo.Echo) *echo.Group {
	App = NewApplication(*NewRepository(map[string]Membership{}))
	membershipGroup := e.Group("/memberships")
	membershipGroup.POST("", Create)
	membershipGroup.PATCH("", Update)
	membershipGroup.DELETE("/:id", Delete)
	membershipGroup.GET("/:id", Read)
	membershipGroup.GET("", ReadAll)
	return membershipGroup
}

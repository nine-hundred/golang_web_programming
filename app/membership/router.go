package membership

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

var App = &Application{}

func SetMembershipRouter(e *echo.Echo, controller Controller) *echo.Echo {
	App = NewApplication(*NewRepository(map[string]Membership{}))
	membershipGroup := e.Group("/memberships")
	membershipGroup.Use(LoggingMiddleware)
	membershipGroup.POST("", controller.Create)
	membershipGroup.PATCH("", controller.Update, OwnerMiddleware)
	membershipGroup.DELETE("/:id", controller.Delete, OwnerMiddleware)
	membershipGroup.GET("/:id", controller.Read, AdminOrOwnerMiddleware)
	membershipGroup.GET("", controller.ReadAll, AdminMiddleware)
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

type MiddlewareResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func AdminOrOwnerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		tokenStr := c.Request().Header.Get("Authorization")
		claims, err := ParseJwt(tokenStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, MiddlewareResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}
		if claims["name"] == "admin" || claims["id"] == id {
			return next(c)
		}
		return c.JSON(http.StatusBadRequest, MiddlewareResponse{
			Code:    http.StatusBadRequest,
			Message: "not admin",
		})
	}

}

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := c.Request().Header.Get("Authorization")
		claims, err := ParseJwt(tokenStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, MiddlewareResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}
		if claims["name"] != "admin" {
			return c.JSON(http.StatusBadRequest, MiddlewareResponse{
				Code:    http.StatusBadRequest,
				Message: "not admin",
			})
		}
		return next(c)
	}
}

func OwnerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		tokenStr := c.Request().Header.Get("Authorization")
		claims, err := ParseJwt(tokenStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, MiddlewareResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}
		if claims["id"] != id {
			return c.JSON(http.StatusBadRequest, MiddlewareResponse{
				Code:    http.StatusBadRequest,
				Message: "not the same id",
			})
		}
		return next(c)
	}
}

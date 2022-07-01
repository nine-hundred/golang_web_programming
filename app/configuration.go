package app

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"golang_web_programming/app/membership"
)

type Config struct {
	Controller membership.Controller
}

func StartValidator() {
	govalidator.TagMap["membershipType"] = govalidator.Validator(func(str string) bool {
		if str == "payco" || str == "naver" || str == "toss" {
			return true
		}
		return false
	})
}

func DefaultConfig() *Config {
	StartValidator()
	data := map[string]membership.Membership{}
	service := membership.NewService(*membership.NewRepository(data))
	controller := membership.NewController(*service)
	return &Config{
		Controller: *controller,
	}
}

func NewEcho(config Config) *echo.Echo {
	e := echo.New()

	controller := config.Controller
	membership.SetMembershipRouter(e, controller)
	e.GET("/logo", controller.GetMemebershipImg)
	e.POST("/login", controller.Login)
	return e
}

package membership

import (
	"crypto/md5"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) Create(c echo.Context) error {
	req := CreateRequest{
		UserName:       c.FormValue("UserName"),
		MembershipType: c.FormValue("MembershipType"),
	}
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return err
	}

	res := controller.service.CreateMembership(req)
	if res.Code != http.StatusCreated {
		return c.JSON(res.Code, res)
	}

	return c.JSON(http.StatusCreated, res)
}

func (controller *Controller) Update(c echo.Context) error {
	req := UpdateRequest{
		ID:             c.FormValue("ID"),
		UserName:       c.FormValue("UserName"),
		MembershipType: c.FormValue("MembershipType"),
	}

	res := controller.service.
		UpdateMembership(req)
	if res.Code != http.StatusOK {
		return c.JSON(res.Code, res)
	}
	return c.JSON(http.StatusOK, res)
}

func (controller Controller) Delete(c echo.Context) error {
	id := c.Param("id")

	res := controller.service.DeleteMembership(id)
	if res.Code != http.StatusOK {
		return c.JSON(res.Code, res)
	}
	return c.JSON(http.StatusOK, res)
}

func (controller Controller) Read(c echo.Context) error {
	id := c.Param("id")

	res := controller.service.ReadMembershipById(id)
	return c.JSON(http.StatusOK, res)
}

func (controller Controller) ReadAll(c echo.Context) error {
	req := ReadAllRequest{
		Limit:  c.QueryParam("limit"),
		Offset: c.QueryParam("offset"),
	}

	res := controller.service.ReadAllMembership(req)
	return c.JSON(http.StatusOK, res)
}

func (controller Controller) GetMemebershipImg(c echo.Context) error {
	url := "app/assets/worldcup.png"

	file, err := os.Stat(url)
	if err != nil {
		return echo.ErrInternalServerError
	}

	modifiedTime := file.ModTime()
	etag := fmt.Sprintf("%x", md5.Sum([]byte(modifiedTime.String())))

	if c.Request().Header.Get("If-None-Match") == etag {
		return c.String(http.StatusNotModified, http.StatusText(http.StatusNotModified))
	}
	c.Response().Header().Set("ETag", etag)
	return c.File(url)
}

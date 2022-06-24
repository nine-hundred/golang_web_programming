package membership

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Create(c echo.Context) error {
	req := CreateRequest{
		UserName:       c.FormValue("UserName"),
		MembershipType: c.FormValue("MembershipType"),
	}

	res, err := App.Create(req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, res)
}

func Update(c echo.Context) error {
	req := UpdateRequest{
		ID:             c.FormValue("ID"),
		UserName:       c.FormValue("UserName"),
		MembershipType: c.FormValue("MembershipType"),
	}

	_, err := App.Update(req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "updated")
}

func Delete(c echo.Context) error {
	id := c.Param("id")

	err := App.Delete(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "deleted")
}

func Read(c echo.Context) error {
	id := c.Param("id")
	res, err := App.Read(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func ReadAll(c echo.Context) error {
	return nil
}

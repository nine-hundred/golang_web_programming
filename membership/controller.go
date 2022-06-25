package membership

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Create(c echo.Context) error {
	req := CreateRequest{
		UserName:       c.FormValue("UserName"),
		MembershipType: c.FormValue("MembershipType"),
	}

	res, err := App.Create(req)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	return c.JSON(http.StatusCreated, res)
}

func Update(c echo.Context) error {
	req := UpdateRequest{
		ID:             c.FormValue("ID"),
		UserName:       c.FormValue("UserName"),
		MembershipType: c.FormValue("MembershipType"),
	}

	res, err := App.Update(req)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	return c.JSON(http.StatusOK, res)
}

func Delete(c echo.Context) error {
	id := c.Param("id")

	err := App.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, DeleteResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, "deleted")
}

func Read(c echo.Context) error {
	id := c.Param("id")
	res, err := App.Read(id)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	return c.JSON(http.StatusOK, res)
}

func ReadAll(c echo.Context) error {
	limit, offset, err := atoiToLimitAndOffset(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ReadResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	req := ReadRequest{
		Limit:  limit,
		Offset: offset,
	}

	res, err := App.ReadAll(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ReadResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, res)
}

func atoiToLimitAndOffset(c echo.Context) (limit int, offset int, err error) {
	if c.QueryParam("limit") == "" || c.QueryParam("offset") == "" {
		return 0, 0, nil
	}
	limit, err = strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return 0, 0, err
	}
	offset, err = strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return 0, 0, err
	}
	return limit, offset, nil
}

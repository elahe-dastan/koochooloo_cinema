package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Favorite struct {
	Store *sql.DB
}

type List struct {
	Film int
	Name string
}

func (f *Favorite) Retrieve(c echo.Context) error {
	name := c.Param("name")

	var list []List
	rows, err := f.Store.Query("SELECT * FROM favorite WHERE name = ?", name)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = rows.Scan(&list); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, list)
}
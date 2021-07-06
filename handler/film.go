package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"koochooloo_cinema/model"
)

type Film struct {
	Store *sql.DB
}

//type filmRequest struct {
//	Tag      string `json:"tag"`
//	Name     string `json:"name"`
//	Producer string `json:"producer"`
//}

func (f *Film) RetrieveByTag(c echo.Context) error {
	tag := c.Param("tag")

	var films []model.Film
	rows, err := f.Store.Query("SELECT * FROM film WHERE tag = ?", tag)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := rows.Scan(&films); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, films)
}

func (f *Film) RetrieveByName(c echo.Context) error {
	tag := c.Param("name")

	var films []model.Film
	rows, err := f.Store.Query("SELECT * FROM film WHERE name = ?", tag)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := rows.Scan(&films); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, films)
}

func (f *Film) RetrieveByProducer(c echo.Context) error {
	tag := c.Param("producer")

	var films []model.Film
	rows, err := f.Store.Query("SELECT * FROM film WHERE producer = ?", tag)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := rows.Scan(&films); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, films)
}
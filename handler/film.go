package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"koochooloo_cinema/request"
)

type Film struct {
	Store *sql.DB
}

type FilmRequest struct {
	Tag   string `query:"tag"`
	Limit int    `query:"name"`
	Page  int    `query:"producer"`
	Ordering string `query:"ordering"`
}

const limit = 10

func (f *Film) RetrieveByTag(c echo.Context) error {
	filmReq := FilmRequest{}
	if err := c.Bind(&filmReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var films []request.Film
	rows, err := f.Store.Query("SELECT * FROM film WHERE tag = ? ORDER BY ? LIMIT ? OFFSET ? ;", filmReq.Tag, filmReq.Ordering, filmReq.Limit, filmReq.Limit*(filmReq.Page-1))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := rows.Scan(&films); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, films)
}

func (f *Film) RetrieveByName(c echo.Context) error {
	filmReq := FilmRequest{}
	if err := c.Bind(&filmReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var films []request.Film
	rows, err := f.Store.Query("SELECT * FROM film WHERE tag = ? ORDER BY ? DESC LIMIT ? OFFSET ? ;", filmReq.Tag, filmReq.Ordering, filmReq.Limit, filmReq.Limit*(filmReq.Page-1))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := rows.Scan(&films); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, films)
}

func (f *Film) RetrieveByProducer(c echo.Context) error {
	filmReq := FilmRequest{}
	if err := c.Bind(&filmReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var films []request.Film
	rows, err := f.Store.Query("SELECT * FROM film WHERE tag = ? ORDER BY ? LIMIT ? OFFSET ? ;", filmReq.Tag, filmReq.Ordering, filmReq.Limit, filmReq.Limit*(filmReq.Page-1))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := rows.Scan(&films); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, films)
}

func (f *Film) Watch(c echo.Context) error {
	film := c.Param("film_id")
	user := c.Param("user_id")

	_, err := f.Store.Exec("INSERT INTO watch VALUES (?, ?)", film, user)
	if err != nil {
		return err // increase by one
	}

	return c.NoContent(http.StatusOK)
}

package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/elahe-dastan/koochooloo_cinema/request"
	"github.com/elahe-dastan/koochooloo_cinema/response"
)

type Film struct {
	Store *sql.DB
}

type FilmRequest struct {
	Tag      string `query:"tag"`
	Limit    int    `query:"name"`
	Page     int    `query:"producer"`
	Ordering string `query:"ordering"`
}

const limit = 10

func (f *Film) RetrieveByTag(c echo.Context) error {
	filmReq := FilmRequest{}
	if err := c.Bind(&filmReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if filmReq.Limit == 0 {
		filmReq.Limit = limit
	}
	if filmReq.Page == 0 {
		filmReq.Page = 1
	}
	if filmReq.Ordering == "" {
		filmReq.Ordering = "id"
	}

	var films []response.Film
	query := fmt.Sprintf("SELECT * FROM film JOIN film_tag ON film.id = film_tag.film WHERE tag = '%s' ORDER BY %s LIMIT %d OFFSET %d ;", filmReq.Tag, filmReq.Ordering, filmReq.Limit, filmReq.Limit*(filmReq.Page-1))
	rows, err := f.Store.Query(query)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	defer rows.Close()

	var ignoreInt int
	var ignoreString string
	for rows.Next() {
		var film response.Film
		// todo what about producers and tags need join
		if err = rows.Scan(&film.ID, &film.File, &film.Name, &film.ProductionYear, &film.Explanation, &film.View, &film.Price, &ignoreInt, &ignoreString); err != nil {
			panic(err)
		}
		films = append(films, film)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, films)
}

func (f *Film) RetrieveById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	query := fmt.Sprintf("SELECT * FROM film WHERE id = %d ;", id)
	row := f.Store.QueryRow(query)

	var film response.Film
	// todo what about producers and tags need join
	if err = row.Scan(&film.ID, &film.File, &film.Name, &film.ProductionYear, &film.Explanation, &film.View, &film.Price); err != nil {
		panic(err)
	}

	if err = row.Err(); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, film)
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

	if err = rows.Scan(&films); err != nil {
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

// nolint: wrapcheck
func (f *Film) Watch(c echo.Context) error {
	user := c.Param("username")

	film, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	query := fmt.Sprintf("INSERT INTO watch VALUES (%d, '%s')", film, user)

	result, err := f.Store.Exec(query)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if n, _ := result.RowsAffected(); n != 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "you must pay for this movie")
	}

	return c.NoContent(http.StatusOK)
}

// Register registers the routes of URL handler on given group.
func (f *Film) Register(g *echo.Group) {
	g.GET("/tag", f.RetrieveByTag)
	g.GET("/film/:id", f.RetrieveById)
	g.GET("/film/:id/watch/:username", f.Watch)
}

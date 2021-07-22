package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/elahe-dastan/koochooloo_cinema/request"
	"github.com/elahe-dastan/koochooloo_cinema/response"

	"github.com/labstack/echo/v4"
)

type Favorite struct {
	Store *sql.DB
}

// nolint: wrapcheck
func (f *Favorite) Create(c echo.Context) error {
	var body request.Favorite
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	query := "INSERT INTO favorite VALUES"

	for i, film := range body.Film {
		if i != 0 {
			query = fmt.Sprintf("%s, ('%s', %d, '%s')", query, body.Username, film, body.Album)
		} else {
			query = fmt.Sprintf("%s ('%s', %d, '%s')", query, body.Username, film, body.Album)
		}
	}

	query += ";"

	result, err := f.Store.Exec(query)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if n, _ := result.RowsAffected(); n != 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "you must be a special user to have a list")
	}

	return c.NoContent(http.StatusCreated)
}

func (f *Favorite) Retrieve(c echo.Context) error {
	username := c.Param("username")
	album := c.Param("album")

	query := fmt.Sprintf("SELECT id, file, name, production_year, explanation, view, price, score FROM favorite JOIN film ON favorite.film = film.id WHERE favorite.username='%s' AND favorite.album='%s';", username, album)

	rows, err := f.Store.Query(query)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	defer rows.Close()

	var films []response.Film
	for rows.Next() {
		var film response.Film
		if err = rows.Scan(&film.ID, &film.File, &film.Name, &film.ProductionYear, &film.Explanation, &film.View, &film.Price, &film.Score); err != nil {
			panic(err)
		}

		films = append(films, film)
	}

	return c.JSON(http.StatusOK, films)
}

func (f *Favorite) Register(g *echo.Group) {
	g.POST("/favorite", f.Create)
	g.GET("/favorite/:username/:album", f.Retrieve)
}

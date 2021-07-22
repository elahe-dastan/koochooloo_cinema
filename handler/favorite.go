package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/elahe-dastan/koochooloo_cinema/request"
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

func (f *Favorite) Register(g *echo.Group) {
	g.POST("/favorite", f.Create)
}

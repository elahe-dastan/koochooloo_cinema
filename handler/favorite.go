package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"koochooloo_cinema/request"
)

type Favorite struct {
	Store *sql.DB
}

func (f *Favorite) Create(c echo.Context) error {
	var body request.Favorite
	err := c.Bind(&body)
	if err != nil {
		return err
	}

	for _, film := range body.Film {
		query := fmt.Sprintf("INSERT INTO favorite VALUES ('%s', %d, '%s')", body.Username, film, body.Album)
		if _, err = f.Store.Exec(query); err != nil {
			return err
		}
	}

	return c.NoContent(http.StatusOK)
}

func (f *Favorite) Register(g *echo.Group) {
	g.POST("/favorite", f.Create)
	//g.POST("/signup", f.Create)
	//g.POST("/edit", f.Update)
	//g.GET("/count/:key", h.Count)
}

package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/elahe-dastan/koochooloo_cinema/request"
)

type Follow struct {
	Store *sql.DB
}

// todo password constraint doesn't work
// todo unique constraint on email doesn't work
func (f *Follow) Create(c echo.Context) error {
	var rq request.Follow
	if err := c.Bind(&rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// todo remove the fmt.Sprintf
	query := fmt.Sprintf("INSERT INTO follow VALUES ('%s', '%s')", rq.Username, rq.Following)
	if _, err := f.Store.Exec(query); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}


// Register registers the routes of URL handler on given group.
func (f *Follow) Register(g *echo.Group) {
	g.POST("/follow", f.Create)
	//g.POST("/signup", f.Create)
	//g.POST("/edit", f.Update)
	//g.GET("/count/:key", h.Count)
}

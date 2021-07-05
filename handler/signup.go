package handler

import (
	"database/sql"
	"net/http"

	"koochooloo_cinema/request"

	"github.com/labstack/echo/v4"
)

type SignUp struct {
	Store  *sql.DB
}

// Create generates short URL and save it on database.
// nolint: wrapcheck
func (s SignUp) Create(c echo.Context) error {
	var rq request.Signup

	if err := c.Bind(&rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if _, err := s.Store.Exec("INSERT INTO registration VALUES (?, ?)", rq.Username, rq.Password); err != nil {
		return err
	}

	// todo return object
	return c.NoContent(http.StatusOK)
}

// Retrieve retrieves URL for given short URL and redirect to it.
// nolint: wrapcheck
func (s SignUp) Retrieve(c echo.Context) error {
	username := c.Param("username")


	user := s.Store.QueryRow("SELECT * FROM registeration WHERE username = ?", username)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.Redirect(http.StatusFound, url)
}


// Register registers the routes of URL handler on given group.
func (s SignUp) Register(g *echo.Group) {
	//g.GET("/:key", h.Retrieve)
	g.POST("/signup", s.Create)
	//g.GET("/count/:key", h.Count)
}
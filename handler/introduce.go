package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/elahe-dastan/koochooloo_cinema/request"
)

type Introduction struct {
	Store *sql.DB
}

func (i *Introduction) Create(c echo.Context) error {
	var introduction request.Introduction
	if err := c.Bind(&introduction); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	query := fmt.Sprintf("INSERT INTO introducer VALUES ('%s', '%s')", introduction.Username, introduction.Introducer)
	if _, err := i.Store.Exec(query); err != nil {
		return err
	}

	// todo return object
	return c.NoContent(http.StatusOK)
}

func (i *Introduction) Retrieve(c echo.Context) error {
	film := c.QueryParam("film")

	var votes request.Vote

	// todo user is keyword?
	//todo what if update
	rows, err := i.Store.Query("SELECT * FROM vote WHERE film = ?", film)
	if err != nil {
		return err
	}

	if err = rows.Scan(votes); err != nil {
		return err
	}

	// todo return object
	return c.JSON(http.StatusOK, votes)
}

func (i *Introduction) Register(g *echo.Group) {
	g.POST("/introduce", i.Create)
	//g.POST("/signup", f.Create)
	//g.POST("/edit", f.Update)
	//g.GET("/count/:key", h.Count)
}
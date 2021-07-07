package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"koochooloo_cinema/model"
)

type Vote struct {
	Store *sql.DB
}

type VoteRequest struct {
	Score   int
	Comment string
}

func (v *Vote) Create(c echo.Context) error {
	user := c.Param("user")
	film := c.Param("film")

	var vote VoteRequest
	if err := c.Bind(&vote); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// todo user is keyword?
	//todo what if update
	if _, err := v.Store.Exec("INSERT INTO vote VALUES (?, ?, ?, ?)", user, film, vote.Score, vote.Comment); err != nil {
		return err
	}

	// todo return object
	return c.NoContent(http.StatusOK)
}

func (v *Vote) Retrieve(c echo.Context) error {
	film := c.QueryParam("film")

	var votes model.Vote

	// todo user is keyword?
	//todo what if update
	rows, err := v.Store.Query("SELECT * FROM vote WHERE film = ?", film)
	if err != nil {
		return err
	}

	if err = rows.Scan(votes); err != nil {
		return err
	}

	// todo return object
	return c.JSON(http.StatusOK, votes)
}
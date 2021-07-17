package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"koochooloo_cinema/request"
)

type Vote struct {
	Store *sql.DB
}

func (v *Vote) Create(c echo.Context) error {
	username := c.Param("username")
	film := c.Param("film")

	var vote request.Vote
	if err := c.Bind(&vote); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//todo what if update
	filmId, err := strconv.Atoi(film)
	if err != nil {
		return err
	}

	tx, err := v.Store.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("SELECT * FROM watch WHERE username = '%s' AND film = %d", username, filmId)
	result, err := v.Store.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if rowsAffected == 0 {
		return c.JSON(http.StatusForbidden, "first watch the movie")
	}

	query = fmt.Sprintf("INSERT INTO vote VALUES ('%s', %d, %d, '%s')", username, filmId, vote.Score, vote.Comment)
	if _, err = v.Store.Exec(query); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	// todo return object
	return c.NoContent(http.StatusOK)
}

func (v *Vote) Retrieve(c echo.Context) error {
	film := c.QueryParam("film")

	var votes request.Vote

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

func (v *Vote) Register(g *echo.Group) {
	g.POST("/comment/:username/:film", v.Create)
	//g.POST("/signup", f.Create)
	//g.POST("/edit", f.Update)
	//g.GET("/count/:key", h.Count)
}
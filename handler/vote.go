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

// nolint: wrapcheck
func (v *Vote) Create(c echo.Context) error {
	username := c.Param("username")
	film := c.Param("film")

	var vote request.Vote
	if err := c.Bind(&vote); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	filmID, err := strconv.Atoi(film)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	query := fmt.Sprintf("INSERT INTO vote VALUES ('%s', %d, %d, '%s')", username, filmID, vote.Score, vote.Comment)

	result, err := v.Store.Exec(query)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if n, _ := result.RowsAffected(); n != 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "you must watch the film before voting")
	}

	return c.NoContent(http.StatusCreated)
}

// nolint: wrapcheck
func (v *Vote) Retrieve(c echo.Context) error {
	film := c.Param("film")

	var votes request.Vote

	rows, err := v.Store.Query("SELECT * FROM vote WHERE film = ?", film)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if rows.Err() != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, rows.Err().Error())
	}
	defer rows.Close()

	if err = rows.Scan(votes); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, votes)
}

func (v *Vote) Register(g *echo.Group) {
	g.POST("/comment/:username/:film", v.Create)
	g.GET("/comment/:film", v.Retrieve)
}

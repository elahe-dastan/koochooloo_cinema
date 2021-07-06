package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

const price = 20

type Special struct {
	Store *sql.DB
}

func (s *Special) Update(c echo.Context) error {
	username := c.Param("username")

	var credit int
	err := s.Store.QueryRow( "SELECT credit FROM wallet WHERE username = ?", username).Scan(&credit)
	if err != nil {
		return err
	}

	if credit < price {
		return c.JSON(http.StatusForbidden, "اول پول وده")
	}

	tx, err := s.Store.Begin()
	if err != nil {
		return err
	}

	// This project is not for production or anything but to handle concurrency we did it
	_, err = tx.Exec( "UPDATE wallet SET credit = credit - ? WHERE username = ?", price, username)
	if err != nil {
		tx.Rollback()
		return err
	}

	// todo set special at
	_, err = tx.Exec( "UPDATE registeration SET credit = credit + 1 WHERE username = ?", username)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Register registers the routes of URL handler on given group.
func (s *Special) Register(g *echo.Group) {
	g.POST("/", s.Update)
}

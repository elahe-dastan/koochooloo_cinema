package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/elahe-dastan/koochooloo_cinema/request"
)

const price = 20

type Special struct {
	Store *sql.DB
}

func (s *Special) UpdateByWallet(c echo.Context) error {
	var body request.Special
	err := c.Bind(&body)
	if err != nil {
		return err
	}

	var credit int
	query := fmt.Sprintf("SELECT credit FROM wallet WHERE username = '%s'", body.Username)
	err = s.Store.QueryRow(query).Scan(&credit)
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
	query = fmt.Sprintf("UPDATE wallet SET credit = credit - %d WHERE username = '%s'", price, body.Username)
	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}

	var specialTill *time.Time
	query = fmt.Sprintf("SELECT special_till FROM users WHERE username = '%s'", body.Username)
	err = tx.QueryRow(query).Scan(&specialTill)
	if err != nil {
		tx.Rollback()
		return err
	}

	//fmt.Println(specialTill.Location())
	//fmt.Println(time.Now().Location())
	fmt.Println(specialTill)
	if specialTill == nil || specialTill.Before(time.Now()) {
		t := time.Now()
		specialTill = &t
	}

	t := specialTill.AddDate(0, 1, 0)
	specialTill = &t
	query = fmt.Sprintf("UPDATE users SET special_till = '%s' WHERE username = '%s'", specialTill.Format(time.RFC3339), body.Username)
	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *Special) UpdateByScore(c echo.Context) error {
	var body request.Special
	err := c.Bind(&body)
	if err != nil {
		return err
	}

	var credit int
	query := fmt.Sprintf("SELECT score FROM users WHERE username = '%s'", body.Username)
	err = s.Store.QueryRow(query).Scan(&credit)
	if err != nil {
		return err
	}

	if credit < 3 {
		return c.JSON(http.StatusForbidden, "you don't have enough score")
	}

	tx, err := s.Store.Begin()
	if err != nil {
		return err
	}

	// This project is not for production or anything but to handle concurrency we did it
	query = fmt.Sprintf("UPDATE users SET score = score - %d WHERE username = '%s'", 1, body.Username)
	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}

	var specialTill *time.Time
	query = fmt.Sprintf("SELECT special_till FROM users WHERE username = '%s'", body.Username)
	err = tx.QueryRow(query).Scan(&specialTill)
	if err != nil {
		tx.Rollback()
		return err
	}

	//fmt.Println(specialTill.Location())
	//fmt.Println(time.Now().Location())
	fmt.Println(specialTill)
	if specialTill == nil || specialTill.Before(time.Now()) {
		t := time.Now()
		specialTill = &t
	}

	t := specialTill.AddDate(0, 1, 0)
	specialTill = &t
	query = fmt.Sprintf("UPDATE users SET special_till = '%s' WHERE username = '%s'", specialTill.Format(time.RFC3339), body.Username)
	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Register registers the routes of URL handler on given group.
func (s *Special) Register(g *echo.Group) {
	g.POST("/special/wallet", s.UpdateByWallet)
	g.POST("/special/score", s.UpdateByScore)
}

package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"koochooloo_cinema/request"

	"github.com/labstack/echo/v4"
)

type SignUp struct {
	Store *sql.DB
}

// todo password constraint doesn't work
// todo unique constraint on email doesn't work
func (s *SignUp) Create(c echo.Context) error {
	var rq request.Signup
	if err := c.Bind(&rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tx, err := s.Store.Begin()
	if err != nil {
		return err
	}

	// todo remove the fmt.Sprintf
	query := fmt.Sprintf("INSERT INTO users (username, password, first_name, last_name, email, phone, national_number) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s')",
		rq.Username, rq.Password, rq.FirstName, rq.LastName, rq.Email, rq.Phone, rq.NationalNumber)
	if _, err = s.Store.Exec(query); err != nil {
		tx.Rollback()
		return err
	}

	// todo what is the error (is the username taken?)
	query = fmt.Sprintf("INSERT INTO wallet (username) VALUES ('%s')", rq.Username)
	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return c.NoContent(http.StatusCreated)
}

// Retrieve retrieves URL for given short URL and redirect to it.
// nolint: wrapcheck
func (s *SignUp) Retrieve(c echo.Context) error {
	username := c.Param("username")

	user := request.Signup{}
	err := s.Store.QueryRow("SELECT * FROM registeration WHERE username = ?", username).Scan(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (s *SignUp) Update(c echo.Context) error {
	var rq request.Signup
	if err := c.Bind(&rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	query := "UPDATE users SET "

	columns := make(map[string]string)

	if rq.Password != "" {
		columns["password"] = rq.Password
	}

	if rq.FirstName != "" {
		columns["first_name"] = rq.FirstName
	}

	if rq.LastName != "" {
		columns["last_name"] = rq.LastName
	}

	if rq.Email != "" {
		columns["email"] = rq.Email
	}

	if rq.Phone != "" {
		columns["phone"] = rq.Phone
	}

	if rq.NationalNumber != "" {
		columns["national_number"] = rq.NationalNumber
	}

	for k, v := range columns {
		query += k + " = '" + v + "', "
	}

	query = strings.Trim(query, ", ")

	query += fmt.Sprintf(" WHERE username = '%s'", rq.Username)

	_, err := s.Store.Exec(query)
	if err != nil {
		return err
	}

	// todo
	//if result.RowsAffected == 0 {
	//	return ctx.JSON(http.StatusNotFound, DriverSignupError{Message: "referrer not found"})
	//}

	return c.NoContent(http.StatusOK)
}

// Register registers the routes of URL handler on given group.
func (s *SignUp) Register(g *echo.Group) {
	g.GET("/:username", s.Retrieve)
	g.POST("/signup", s.Create)
	g.POST("/edit", s.Update)
	//g.GET("/count/:key", h.Count)
}

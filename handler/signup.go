package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"koochooloo_cinema/model"
	"koochooloo_cinema/request"

	"github.com/labstack/echo/v4"
)

type SignUp struct {
	Store *sql.DB
}

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

	user := model.User{}
	err := s.Store.QueryRow("SELECT * FROM registeration WHERE username = ?", username).Scan(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (s *SignUp) Update(c echo.Context) error {
	username := c.Param("username")

	body := model.User{}
	err := c.Bind(&body)
	if err != nil {
		return err
	}

	query := "UPDATE registeration SET "

	columns := make(map[string]string)

	if body.Username != "" {
		columns["username"] = body.Username
	}

	if body.Password != "" {
		columns["password"] = body.Password
	}

	if body.FirstName != "" {
		columns["first_name"] = body.FirstName
	}

	if body.LastName != "" {
		columns["last_name"] = body.LastName
	}

	if body.Email != "" {
		columns["email"] = body.Email
	}

	if body.Phone != "" {
		columns["phone"] = body.Phone
	}

	if body.NationalNumber != "" {
		columns["national_number"] = body.NationalNumber
	}

	for k, v := range columns {
		query += k + " = " + v + ", "
	}

	query = strings.Trim(query, ",")

	query += fmt.Sprintf("WHERE username = %s", username)

	_, err = s.Store.Query(query)
	if err != nil {
		return err
	}

	//if result.RowsAffected == 0 {
	//	return ctx.JSON(http.StatusNotFound, DriverSignupError{Message: "referrer not found"})
	//}

	//return ctx.JSON(http.StatusOK, &ReferrerResponse{
	//	Name:            referrer.Name,
	//	Code:            referrer.Code,
	//	CreatedAt:       referrer.CreatedAt,
	//	UpdatedAt:       referrer.UpdatedAt,
	//	Status:          status,
	//	UploadPermitted: &referrer.UploadPermitted,
	//	Email:           referrer.Email,
	//	Cellphone:       referrer.Cellphone,
	//})

	return c.NoContent(http.StatusOK)
	// todo
	//تغییر موجودی حساب کاربری و نام کاربری نباید
	//در این قسمت امکان پذیر باشد و در صورت تغییر باید تمام تغییرات rollback شوند
}

// Register registers the routes of URL handler on given group.
func (s *SignUp) Register(g *echo.Group) {
	g.GET("/:username", s.Retrieve)
	g.POST("/signup", s.Create)
	//g.GET("/count/:key", h.Count)
}

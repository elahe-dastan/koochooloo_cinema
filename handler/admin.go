package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"koochooloo_cinema/request"
)

// todo
const (
	username = "admin"
	password = "admin"
)

type Admin struct {
	Store *sql.DB
}

func (a *Admin) Create(c echo.Context) error {
	var film request.Film

	if err := c.Bind(&film); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tx, err := a.Store.Begin()
	if err != nil {
		return err
	}

	id := 0
	query := fmt.Sprintf("INSERT INTO film (file,  name, production_year, explanation, price) VALUES ('%s', '%s', %d, '%s', %d) RETURNING id",
		film.File, film.Name, film.ProductionYear, film.Explanation, film.Price)
	err = a.Store.QueryRow(query).Scan(&id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// todo this is so bad
	if film.View != 0 {
		query = fmt.Sprintf("UPDATE film SET view = %d WHERE id = %d",film.View, id)
		_, err = a.Store.Exec(query)
		if err != nil {
			tx.Rollback()
			return err
		}

	}

	// todo this is so bad
	if film.Price != 0 {
		query = fmt.Sprintf("UPDATE film SET price = %d WHERE id = %d",film.Price, id)
		_, err = a.Store.Exec(query)
		if err != nil {
			tx.Rollback()
			return err
		}

	}

	for _, producer := range film.Producers {
		query = fmt.Sprintf("INSERT INTO film_producer (film, producer) VALUES ('%d', '%s')", id, producer)
		if _, err = a.Store.Exec(query); err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, tag := range film.Tags {
		query = fmt.Sprintf("INSERT INTO film_tag (film, tag) VALUES ('%d', '%s')", id, tag)
		if _, err = a.Store.Exec(query); err != nil {
			tx.Rollback()
			return err
		}
	}

	// todo return object
	return c.NoContent(http.StatusOK)
}

// Retrieve retrieves URL for given short URL and redirect to it.
// nolint: wrapcheck
//todo tag
func (a *Admin) Retrieve(c echo.Context) error {
	username := c.Param("username")

	user := request.Signup{}
	err := a.Store.QueryRow("SELECT * FROM registeration WHERE username = ?", username).Scan(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (a *Admin) Update(c echo.Context) error {
	username := c.Param("username")

	body := request.Signup{}
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

	_, err = a.Store.Query(query)
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

func (a *Admin) Delete(c echo.Context) error {
	return nil
}

// Register registers the routes of URL handler on given group.
func (a *Admin) Register(g *echo.Group) {
	g.GET("/:username", a.Retrieve)
	g.POST("/admin", a.Create)
	//g.GET("/count/:key", h.Count)
}

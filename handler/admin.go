package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"koochooloo_cinema/model"
)

const (
	username = "admin"
	password = "admin"
)

type Admin struct {
	Store *sql.DB
}

// Create generates short URL and save it on database.
// nolint: wrapcheck
func (a *Admin) Create(c echo.Context) error {
	var film model.Film

	if err := c.Bind(&film); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//todo others
	if _, err := a.Store.Exec("INSERT INTO film VALUES (?, ?)", film.File, film.Explanation); err != nil {
		return err
	}

	// todo return object
	return c.NoContent(http.StatusOK)
}

// Retrieve retrieves URL for given short URL and redirect to it.
// nolint: wrapcheck
//todo tag
func (a *Admin) Retrieve(c echo.Context) error {
	username := c.Param("username")

	user := model.User{}
	err := a.Store.QueryRow("SELECT * FROM registeration WHERE username = ?", username).Scan(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (a *Admin) Update(c echo.Context) error {
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
	g.POST("/signup", a.Create)
	//g.GET("/count/:key", h.Count)
}

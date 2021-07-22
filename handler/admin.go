package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/elahe-dastan/koochooloo_cinema/request"
	"github.com/labstack/echo/v4"
)

// todo
const (
//username = "admin"
//password = "admin"
)

type Admin struct {
	Store *sql.DB
}

// nolint: wrapcheck
func (a *Admin) Create(c echo.Context) error {
	var film request.Film
	if err := c.Bind(&film); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tx, err := a.Store.Begin()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	id := 0
	query := fmt.Sprintf("INSERT INTO film (file,  name, production_year, explanation, price) VALUES ('%s', '%s', %d, '%s', %d) RETURNING id",
		film.File, film.Name, film.ProductionYear, film.Explanation, film.Price)

	if err := a.Store.QueryRow(query).Scan(&id); err != nil {
		_ = tx.Rollback()

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if film.View != 0 {
		query = fmt.Sprintf("UPDATE film SET view = %d WHERE id = %d", film.View, id)

		_, err = a.Store.Exec(query)
		if err != nil {
			_ = tx.Rollback()

			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	if film.Price != 0 {
		query = fmt.Sprintf("UPDATE film SET price = %d WHERE id = %d", film.Price, id)

		_, err = a.Store.Exec(query)
		if err != nil {
			_ = tx.Rollback()

			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	for _, producer := range film.Producers {
		query = fmt.Sprintf("INSERT INTO film_producer (film, producer) VALUES ('%d', '%s')", id, producer)
		if _, err = a.Store.Exec(query); err != nil {
			_ = tx.Rollback()

			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	for _, tag := range film.Tags {
		query = fmt.Sprintf("INSERT INTO film_tag (film, tag) VALUES ('%d', '%s')", id, tag)
		if _, err = a.Store.Exec(query); err != nil {
			_ = tx.Rollback()

			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.NoContent(http.StatusCreated)
}

// Retrieve retrieves URL for given short URL and redirect to it.
// nolint: wrapcheck
//func (a *Admin) Retrieve(c echo.Context) error {
//	username := c.Param("username")
//
//	user := request.Signup{}
//	err := a.Store.QueryRow("SELECT * FROM registeration WHERE username = ?", username).Scan(&user)
//	if err != nil {
//		return echo.NewHTTPError(http.StatusNotFound, err.Error())
//	}
//
//	return c.JSON(http.StatusOK, user)
//}

func (a *Admin) Update(c echo.Context) error {
	id := c.Param("id")
	var body request.Film
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	query := "UPDATE film SET "

	columns := make(map[string]string)

	if body.File != "" {
		columns["file"] = body.File
	}

	if body.Name != "" {
		columns["name"] = body.Name
	}

	if body.Explanation != "" {
		columns["explanation"] = body.Explanation
	}

	if body.Price != 0 {
		columns["price"] = strconv.Itoa(body.Price)
	}

	if body.ProductionYear != 0 {
		columns["production_year"] = strconv.Itoa(body.ProductionYear)
	}

	if body.View != 0 {
		columns["view"] = strconv.Itoa(body.View)
	}

	for k, v := range columns {
		query += k + " = '" + v + "', "
	}

	query = strings.Trim(query, ", ")

	query += fmt.Sprintf(" WHERE id = '%s'", id)

	_, err := a.Store.Query(query)
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
	//g.GET("/:username", a.Retrieve)
	g.POST("/admin", a.Create)
	g.POST("/admin/update/:id", a.Update)
	//g.GET("/count/:key", h.Count)
}

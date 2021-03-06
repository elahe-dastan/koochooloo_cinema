package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/elahe-dastan/koochooloo_cinema/request"
	"github.com/elahe-dastan/koochooloo_cinema/response"
)

type Film struct {
	Store *sql.DB
}

type FilmRequest struct {
	Tag      string `query:"tag"`
	Name     string `query:"name"`
	Producer string `query:"producer"`
	Limit    int    `query:"limit"`
	Page     int    `query:"page"`
	Ordering string `query:"ordering"`
	Special  bool   `query:"special"`
}

const limit = 10

// nolint: wrapcheck
func (f *Film) Retrieve(c echo.Context) error {
	var req FilmRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if req.Limit == 0 {
		req.Limit = limit
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if req.Ordering == "" {
		req.Ordering = "id"
	}

	var films []response.Film
	query := "SELECT id, file, name, production_year, explanation, view, price, score FROM film "

	if req.Special {
		query += "WHERE price > 0"
	}

	query += fmt.Sprintf("ORDER BY %s LIMIT %d OFFSET %d ;",
		req.Ordering,
		req.Limit,
		req.Limit*(req.Page-1),
	)

	rows, err := f.Store.Query(query)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var film response.Film
		if err = rows.Scan(&film.ID, &film.File, &film.Name, &film.ProductionYear, &film.Explanation, &film.View, &film.Price, &film.Score); err != nil {
			panic(err)
		}

		queryTag := fmt.Sprintf("SELECT tag FROM film_tag WHERE film=%d", film.ID)
		rowsTag, err := f.Store.Query(queryTag)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		for rowsTag.Next() {
			var tag string
			if err = rowsTag.Scan(&tag); err != nil {
				panic(err)
			}

			film.Tags = append(film.Tags, tag)
		}

		queryProducer := fmt.Sprintf("SELECT producer FROM film_producer WHERE film=%d", film.ID)
		rowsProducer, err := f.Store.Query(queryProducer)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		for rowsProducer.Next() {
			var producer string
			if err = rowsProducer.Scan(&producer); err != nil {
				panic(err)
			}

			film.Producers = append(film.Producers, producer)
		}

		films = append(films, film)
	}

	return c.JSON(http.StatusOK, films)
}

func (f *Film) RetrieveByTag(c echo.Context) error {
	filmReq := FilmRequest{}
	if err := c.Bind(&filmReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if filmReq.Limit == 0 {
		filmReq.Limit = limit
	}
	if filmReq.Page == 0 {
		filmReq.Page = 1
	}
	if filmReq.Ordering == "" {
		filmReq.Ordering = "id"
	}

	var films []response.Film
	query := fmt.Sprintf("SELECT * FROM film JOIN film_tag ON film.id = film_tag.film WHERE tag = '%s' ORDER BY %s LIMIT %d OFFSET %d ;", filmReq.Tag, filmReq.Ordering, filmReq.Limit, filmReq.Limit*(filmReq.Page-1))
	rows, err := f.Store.Query(query)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	defer rows.Close()

	var ignoreInt int
	var ignoreString string
	for rows.Next() {
		var film response.Film
		// todo what about producers and tags need join
		if err = rows.Scan(&film.ID, &film.File, &film.Name, &film.ProductionYear, &film.Explanation, &film.View, &film.Price, &ignoreInt, &ignoreString); err != nil {
			panic(err)
		}
		films = append(films, film)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, films)
}

// nolint: wrapcheck
func (f *Film) RetrieveByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	query := fmt.Sprintf("SELECT * FROM film WHERE id = %d ;", id)
	row := f.Store.QueryRow(query)

	var film response.Film
	if err = row.Scan(&film.ID, &film.File, &film.Name, &film.ProductionYear,
		&film.Explanation, &film.View, &film.Price, &film.Score); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err = row.Err(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, film)
}

func (f *Film) RetrieveByName(c echo.Context) error {
	filmReq := FilmRequest{}
	if err := c.Bind(&filmReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if filmReq.Limit == 0 {
		filmReq.Limit = limit
	}
	if filmReq.Page == 0 {
		filmReq.Page = 1
	}
	if filmReq.Ordering == "" {
		filmReq.Ordering = "id"
	}

	var films []response.Film
	query := fmt.Sprintf("SELECT * FROM film WHERE name = '%s' ORDER BY %s DESC LIMIT %d OFFSET %d ;", filmReq.Name, filmReq.Ordering, filmReq.Limit, filmReq.Limit*(filmReq.Page-1))
	rows, err := f.Store.Query(query)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var film response.Film
		// todo what about producers and tags need join
		if err = rows.Scan(&film.ID, &film.File, &film.Name, &film.ProductionYear, &film.Explanation, &film.View, &film.Price, &film.Score); err != nil {
			panic(err)
		}
		films = append(films, film)
	}

	return c.JSON(http.StatusOK, films)
}

func (f *Film) RetrieveByProducer(c echo.Context) error {
	filmReq := FilmRequest{}
	if err := c.Bind(&filmReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if filmReq.Limit == 0 {
		filmReq.Limit = limit
	}
	if filmReq.Page == 0 {
		filmReq.Page = 1
	}
	if filmReq.Ordering == "" {
		filmReq.Ordering = "id"
	}

	var films []request.Film
	query := fmt.Sprintf("SELECT * FROM film JOIN film_producer ON film.id = film_producer.film WHERE producer = '%s' ORDER BY %s LIMIT %d OFFSET %d ;", filmReq.Producer, filmReq.Ordering, filmReq.Limit, filmReq.Limit*(filmReq.Page-1))
	rows, err := f.Store.Query(query)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := rows.Scan(&films); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, films)
}

// nolint: wrapcheck
func (f *Film) Watch(c echo.Context) error {
	user := c.Param("username")

	film, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	query := fmt.Sprintf("INSERT INTO watch VALUES (%d, '%s', 1);", film, user)

	result, err := f.Store.Exec(query)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if n, _ := result.RowsAffected(); n != 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "you must pay for this movie")
	}

	return c.NoContent(http.StatusOK)
}

func (f *Film) WatchByScore(c echo.Context) error {
	user := c.Param("username")

	film, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tx, err := f.Store.Begin()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	query := fmt.Sprintf("INSERT INTO watch_score VALUES (%d, '%s', 1);", film, user)

	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	query = fmt.Sprintf("UPDATE users SET score = score - 1 WHERE username = '%s'", user)

	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	tx.Commit()

	return c.NoContent(http.StatusOK)
}

// Register registers the routes of URL handler on given group.
func (f *Film) Register(g *echo.Group) {
	g.GET("/film", f.Retrieve)
	g.GET("/tag", f.RetrieveByTag)
	g.GET("/name", f.RetrieveByName)
	g.GET("/producer", f.RetrieveByProducer)
	g.GET("/film/:id", f.RetrieveByID)
	g.GET("/film/:id/watch/:username", f.Watch)
	g.GET("/film/watch/score/:id/:username", f.WatchByScore)
}

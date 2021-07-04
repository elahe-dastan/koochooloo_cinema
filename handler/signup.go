package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"koochooloo_cinema/request"
)

type SignUp struct {
	Store  *sql.DB
}

// Create generates short URL and save it on database.
// nolint: wrapcheck
func (s SignUp) Create(c echo.Context) error {
	var rq request.Signup

	if err := c.Bind(&rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if _, err := s.Store.Exec("INSERT INTO registration VALUES (?, ?)", rq.Username, rq.Password); err != nil {
		return err
	}

	// todo return object
	return c.NoContent(http.StatusOK)
}

// Retrieve retrieves URL for given short URL and redirect to it.
// nolint: wrapcheck
func (h URL) Retrieve(c echo.Context) error {
	ctx, span := h.Tracer.Start(c.Request().Context(), "handler.url.retrieve")
	defer span.End()

	key := c.Param("key")

	url, err := h.Store.Get(ctx, key)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := h.Store.Inc(ctx, key); err != nil {
		h.Logger.Error("increase counter for fetching url failed",
			zap.Error(err),
			zap.String("key", key),
			zap.String("url", url),
		)
	}

	return c.Redirect(http.StatusFound, url)
}

// Count retrieves the access count for the given short URL.
// nolint: wrapcheck
func (h URL) Count(c echo.Context) error {
	ctx, span := h.Tracer.Start(c.Request().Context(), "handler.url.count")
	defer span.End()

	key := c.Param("key")

	count, err := h.Store.Count(ctx, key)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, count)
}

// Register registers the routes of URL handler on given group.
func (s SignUp) Register(g *echo.Group) {
	//g.GET("/:key", h.Retrieve)
	g.POST("/signup", s.Create)
	//g.GET("/count/:key", h.Count)
}
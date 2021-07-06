package handler

import (
	"database/sql"
	"net/http"

	"koochooloo_cinema/model"

	"github.com/labstack/echo/v4"
)

type Wallet struct {
	Store *sql.DB
}

func (w *Wallet) Update(c echo.Context) error {
	body := model.Wallet{}
	err := c.Bind(&body)
	if err != nil {
		return err
	}

	_, err = w.Store.Query( "UPDATE wallet SET credit = ? WHERE username = ?", body.Credit, body.Username)
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
func (w *Wallet) Register(g *echo.Group) {
	g.POST("/", w.Update)
}

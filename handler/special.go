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
	//د. در صورت وجود رکورد مرتبط به کاربر در جدول کاربران ویژه،
	//اگر زمان اشتراک کاربر تمام شده، تا یک ماه بعد تمدید شود و در صورت وجود اشتراک، زمان اشتراک کاربر به
	//مدت 1 ماه اضافه شود.
		_, err = tx.Exec( "UPDATE registeration SET credit = credit + 1 WHERE username = ?", username)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *Special) UpdateByScore(c echo.Context) error {
	user := c.Param("user")

	var score int
	err := s.Store.QueryRow( "SELECT score FROM registeration WHERE username = ?", user).Scan(&score)
	if err != nil {
		return err
	}

	if score < 3 {
		return c.JSON(http.StatusForbidden, "increase your score")
	}

	tx, err := s.Store.Begin()
	if err != nil {
		return err
	}

	// This project is not for production or anything but to handle concurrency we did it
	_, err = tx.Exec( "UPDATE registeration SET score = score - ? WHERE user = ?", 3, user)
	if err != nil {
		tx.Rollback()
		return err
	}

	// todo set special at
	//د. در صورت وجود رکورد مرتبط به کاربر در جدول کاربران ویژه،
	//اگر زمان اشتراک کاربر تمام شده، تا یک ماه بعد تمدید شود و در صورت وجود اشتراک، زمان اشتراک کاربر به
	//مدت 1 ماه اضافه شود.
	_, err = tx.Exec( "UPDATE registeration SET credit = credit + 1 WHERE username = ?", user)
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

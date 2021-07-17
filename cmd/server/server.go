package server

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"

	"koochooloo_cinema/db"
	"koochooloo_cinema/handler"
)

func main() {
	app := echo.New()

	database, err := db.New()
	if err != nil {
		log.Fatal("database initiation failed", err)
	}

	signup := handler.SignUp{
		Store: database,
	}
	signup.Register(app.Group("/api"))

	wallet := handler.Wallet{
		Store: database,
	}
	wallet.Register(app.Group("/api"))

	special := handler.Special{
		Store: database,
	}
	special.Register(app.Group("/api"))

	admin := handler.Admin{
		Store: database,
	}
	admin.Register(app.Group("/api"))

	film := handler.Film{
		Store: database,
	}
	film.Register(app.Group("/api"))

	vote := handler.Vote{
		Store: database,
	}
	vote.Register(app.Group("/api"))

	if err = app.Start(":1373"); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("echo initiation failed", err)
	}
}

// Register server command.
func Register(root *cobra.Command) {
	root.AddCommand(
		// nolint: exhaustivestruct
		&cobra.Command{
			Use:   "serve",
			Short: "Run server to serve the requests",
			Run: func(cmd *cobra.Command, args []string) {
				main()
			},
		},
	)
}

package migrate

import (
	"koochooloo_cinema/db"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

func main() {
	database, err := db.New()
	if err != nil {
		log.Fatal("database initiation failed", err)
	}

	driver, err := postgres.WithInstance(database, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migration",
		"postgres", driver)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if err = m.Up(); err != nil {
		log.Fatalf(err.Error())
	}
}

// Register migrate command.
func Register(root *cobra.Command) {
	root.AddCommand(
		// nolint: exhaustivestruct
		&cobra.Command{
			Use:   "migrate",
			Short: "Setup database indices",
			Run: func(cmd *cobra.Command, args []string) {
				main()
			},
		},
	)
}

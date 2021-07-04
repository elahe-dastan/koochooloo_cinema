package migrate

import (
	"koochooloo_cinema/db"

	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

const enable = 1

func main() {
	database, err := db.New()
	if err != nil {
		log.Fatal("database initiation failed", err)
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
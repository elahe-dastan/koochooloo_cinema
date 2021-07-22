package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/elahe-dastan/koochooloo_cinema/cmd/migrate"
	"github.com/elahe-dastan/koochooloo_cinema/cmd/server"
)

// ExitFailure status code.
const ExitFailure = 1

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	//cfg := config.New()

	// nolint: exhaustivestruct
	root := &cobra.Command{
		Use:   "github.com/elahe-dastan/koochooloo_cinema",
		Short: "github.com/elahe-dastan/koochooloo_cinema",
	}

	server.Register(root)
	migrate.Register(root)

	if err := root.Execute(); err != nil {
		os.Exit(ExitFailure)
	}
}

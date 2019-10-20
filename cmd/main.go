package main

import (
	"github.com/anfimovoleh/blockchain-sample"
	"github.com/anfimovoleh/blockchain-sample/conf"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func main() {
	cfg := conf.New()

	root := cobra.Command{}
	start := &cobra.Command{
		Use:     "start",
		Short:   "start Blockchain Node",
		Long:    "Start the Blockchain Node with provided REST'ful API and empty ledger",
		Example: "./node start",
		Run: func(cmd *cobra.Command, args []string) {
			app := app.New(cfg)

			if err := app.Start(); err != nil {
				panic(errors.Wrap(err, "failed to start app"))
			}
		},
	}

	root.AddCommand(start)
	if err := root.Execute(); err != nil {
		panic(err)
	}
}

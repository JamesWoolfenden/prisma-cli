package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"os"
	prisma "prisma-cli/src"
)

func main() {
	var apiKey string
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "suppression",
				Aliases: []string{"s"},
				Usage:   "iac code suppressions",
				Subcommands: []*cli.Command{
					{
						Name:    "export",
						Usage:   "export suppression",
						Aliases: []string{"x"},
						Action: func(cCtx *cli.Context) error {
							return prisma.Export(apiKey)
						},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "key",
								Aliases:     []string{"k"},
								Destination: &apiKey,
							},
						},
					},
					{
						Name:  "import",
						Usage: "import suppressions",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("removed task template: ", cCtx.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err)
	}
}

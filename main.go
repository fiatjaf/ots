package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"
)

var log = zerolog.New(os.Stderr).Output(zerolog.ConsoleWriter{Out: os.Stderr})

func main() {
	app := &cli.App{
		Name:  "ots",
		Usage: "a simple opentimestamps cli tool",
		Commands: []*cli.Command{
			stamp,
			upgrade,
			info,
			verify,
			clean,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

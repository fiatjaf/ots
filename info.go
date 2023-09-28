package main

import (
	"fmt"
	"os"

	"github.com/fiatjaf/opentimestamps"
	"github.com/urfave/cli/v2"
)

var info = &cli.Command{
	Name:        "info",
	Usage:       "",
	Description: `reads an .ots file and displays its contents.`,
	Flags:       []cli.Flag{},
	ArgsUsage:   "[file]",
	Action: func(c *cli.Context) error {
		file := c.Args().First()
		b, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}

		data, err := opentimestamps.ReadFromFile(b)
		if err != nil {
			panic(err)
		}

		fmt.Println(data.Human())
		return nil
	},
}

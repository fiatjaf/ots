package main

import (
	"fmt"
	"os"

	"github.com/fiatjaf/opentimestamps"
	"github.com/urfave/cli/v2"
)

var verify = &cli.Command{
	Name:        "verify",
	Usage:       "",
	Description: ``,
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

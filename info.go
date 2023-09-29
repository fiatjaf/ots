package main

import (
	"fmt"
	"os"

	"github.com/nbd-wtf/opentimestamps"
	"github.com/urfave/cli/v2"
)

var info = &cli.Command{
	Name:        "info",
	Description: `reads an .ots file and displays its contents in a readable way.`,
	Flags:       []cli.Flag{},
	ArgsUsage:   "[file]",
	Action: func(c *cli.Context) error {
		file := c.Args().First()
		b, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		data, err := opentimestamps.ReadFromFile(b)
		if err != nil {
			return err
		}

		fmt.Println(data.Human())
		return nil
	},
}

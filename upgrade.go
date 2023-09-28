package main

import (
	"context"
	"fmt"
	"os"

	"github.com/fiatjaf/opentimestamps"
	"github.com/urfave/cli/v2"
)

var upgrade = &cli.Command{
	Name:        "upgrade",
	Usage:       "",
	Description: `reads an .ots file and tries to upgrade it against its specified pending calendars.`,
	Flags:       []cli.Flag{},
	ArgsUsage:   "[file]",
	Action: func(c *cli.Context) error {
		file := c.Args().First()
		b, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		// fmt.Println("file", hex.EncodeToString(b))
		ts, err := opentimestamps.ReadFromFile(b)
		if err != nil {
			return err
		}

		for _, seq := range ts.GetPendingSequences() {
			_, err := seq.Upgrade(context.Background(), ts.Digest)
			if err != nil {
				return err
			}

			fmt.Println("bingo")
		}

		return nil
	},
}

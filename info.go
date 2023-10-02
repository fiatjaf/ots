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
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "onlyfinal",
			Aliases: []string{"f"},
			Usage:   "filter out all pending sequences, leaving only Bitcoin-attested sequences",
		},
		&cli.BoolFlag{
			Name:    "oldest",
			Aliases: []string{"o"},
			Usage:   "leave only the oldest Bitcoin block attestation",
		},
		&cli.BoolFlag{
			Name:    "onlypending",
			Aliases: []string{"p"},
			Usage:   "filter out all Bitcoin-attested sequences, leaving only pending sequences",
		},
	},
	ArgsUsage: "[file]",
	Action: func(c *cli.Context) error {
		file := c.Args().First()
		b, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		ts, err := opentimestamps.ReadFromFile(b)
		if err != nil {
			return err
		}

		if c.Bool("onlyfinal") {
			ts.Sequences = ts.GetBitcoinAttestedSequences()
		} else if c.Bool("onlypending") {
			ts.Sequences = ts.GetPendingSequences()
		}

		if c.Bool("oldest") {
			sequences := []opentimestamps.Sequence{}
			oldest := 0 // we'll invert everything for brevity
			for _, seq := range ts.GetBitcoinAttestedSequences() {
				if bbh := seq.GetAttestation().BitcoinBlockHeight; -1*int(bbh) < oldest {
					oldest = int(bbh)
					sequences = []opentimestamps.Sequence{seq}
				}
			}
			ts.Sequences = sequences
		}

		fmt.Println(ts.Human())
		return nil
	},
}

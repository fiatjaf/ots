package main

import (
	"fmt"
	"os"

	"github.com/nbd-wtf/opentimestamps"
	"github.com/urfave/cli/v2"
)

var clean = &cli.Command{
	Name:        "clean",
	Description: `modifies a file to leave only the oldest Bitcoin attestation`,
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

		var oldest opentimestamps.Sequence
		for _, seq := range ts.GetBitcoinAttestedSequences() {
			if len(oldest) == 0 || seq.GetAttestation().BitcoinBlockHeight < oldest.GetAttestation().BitcoinBlockHeight {
				oldest = seq
			}
		}

		if len(oldest) < 0 {
			return fmt.Errorf("no bitcoin attested sequences found")
		}
		if len(ts.Sequences) == 1 {
			fmt.Fprintf(os.Stderr, "already clear")
			return nil
		}

		ts.Sequences = []opentimestamps.Sequence{oldest}

		if err := os.Rename(file, file+".full"); err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "renamed '%s' to '%s'\n", file, file+".full")

		if err := os.WriteFile(file, ts.SerializeToFile(), 0644); err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "saved new file '%s'\n", file)

		return nil
	},
}

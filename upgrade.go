package main

import (
	"context"
	"fmt"
	"os"

	"github.com/nbd-wtf/opentimestamps"
	"github.com/urfave/cli/v2"
)

var upgrade = &cli.Command{
	Name:        "upgrade",
	Description: `reads an .ots file and tries to upgrade it against its specified pending calendars.`,
	Flags:       []cli.Flag{},
	ArgsUsage:   "[file]",
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

		upgraded := false
		for _, seq := range ts.GetPendingSequences() {
			tail, err := seq.Upgrade(context.Background(), ts.Digest)
			if err != nil {
				fmt.Fprintf(os.Stderr, "- upgrade failed: %s\n", err)
				continue
			}

			newSeq := make(opentimestamps.Sequence, len(seq)+len(tail)-1)
			copy(newSeq, seq[0:len(seq)-1])
			copy(newSeq[len(seq)-1:], tail)

			ts.Sequences = append(ts.Sequences, newSeq)
			fmt.Fprintf(os.Stderr, "- upgraded sequence on %s to bitcoin block %d\n",
				seq[len(seq)-1].Attestation.CalendarServerURL, newSeq[len(newSeq)-1].Attestation.BitcoinBlockHeight)
			upgraded = true
		}

		if !upgraded {
			return fmt.Errorf("unable to upgrade '%s'", file)
		}

		if err := os.Rename(file, file+".bak"); err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "renamed '%s' to '%s'\n", file, file+".bak")

		if err := os.WriteFile(file, ts.SerializeToFile(), 0644); err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "saved new file '%s'\n", file)

		return nil
	},
}

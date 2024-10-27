package main

import (
	"context"
	"fmt"
	"os"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/nbd-wtf/opentimestamps"
	"github.com/urfave/cli/v3"
)

const bitcoinCategory = "options for using a local bitcoind node through RPC:"

var verify = &cli.Command{
	Name:        "verify",
	Description: `verifies all the bitcoin proofs in an .ots file against a bitcoind node or esplora http server.`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "bitcoinrpc-host",
			DefaultText: "127.0.0.1",
			Value:       "127.0.0.1",
			Category:    bitcoinCategory,
		},
		&cli.StringFlag{
			Name:        "bitcoinrpc-port",
			DefaultText: "8332",
			Value:       "8332",
			Category:    bitcoinCategory,
		},
		&cli.StringFlag{
			Name:     "bitcoinrpc-user",
			Category: bitcoinCategory,
		},
		&cli.StringFlag{
			Name:     "bitcoinrpc-password",
			Category: bitcoinCategory,
		},
		&cli.StringFlag{
			Name:        "esplora",
			Usage:       "internet base address of an esplora API",
			DefaultText: "https://blockstream.info/api",
			Value:       "https://blockstream.info/api",
		},
	},
	ArgsUsage: "[file]",
	Action: func(ctx context.Context, c *cli.Command) error {
		b, err := readFromStdinOrFile(c)
		if err != nil {
			return err
		}

		ts, err := opentimestamps.ReadFromFile(b)
		if err != nil {
			return err
		}

		var bitcoin opentimestamps.Bitcoin
		bitcoind := rpcclient.ConnConfig{
			Host:         c.String("bitcoinrpc-host") + ":" + c.String("bitcoinrpc-port"),
			User:         c.String("bitcoinrpc-user"),
			Pass:         c.String("bitcoinrpc-password"),
			HTTPPostMode: true,
		}
		if bitcoind.User != "" && bitcoind.Pass != "" {
			fmt.Fprintf(os.Stderr, "> using a bitcoind node at %s\n", bitcoind.Host)
			bitcoin, err = opentimestamps.NewBitcoindInterface(bitcoind)
			if err != nil {
				return fmt.Errorf("error trying to make a bitcoind connection: %w", err)
			}
		} else if esplora := c.String("esplora"); esplora != "" {
			fmt.Fprintf(os.Stderr, "> using an esplora server at %s\n", esplora)
			bitcoin = opentimestamps.NewEsploraClient(esplora)
		} else {
			return fmt.Errorf("need a way to inspect the bitcoin blockchain")
		}

		anySequences := 0
		valid := make([]uint64, 0, 5)
		for _, seq := range ts.GetBitcoinAttestedSequences() {
			anySequences++
			block := seq[len(seq)-1].Attestation.BitcoinBlockHeight

			if tx, err := seq.Verify(bitcoin, ts.Digest); err != nil {
				fmt.Fprintf(os.Stderr, "- sequence ending on block %d is invalid: %s\n", block, err)
			} else {
				fmt.Fprintf(os.Stderr, "- sequence ending on block %d is valid, tx: %s\n", block, tx.TxHash())
				valid = append(valid, block)
			}
		}

		if len(valid) == 0 {
			if anySequences == 0 {
				return fmt.Errorf("no bitcoin attestations found")
			}

			return fmt.Errorf("no valid sequences found")
		} else {
			s := "s"
			if len(valid) == 1 {
				s = ""
			}
			fmt.Printf("timestamp validated at block%s %v\n", s, valid)
			return nil
		}
	},
}

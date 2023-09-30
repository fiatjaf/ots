package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nbd-wtf/opentimestamps"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
)

var stamp = &cli.Command{
	Name:        "stamp",
	Description: `creates a timestamp from a file or from a 32-byte hash using the given calendar servers.`,
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:        "calendar",
			Aliases:     []string{"c"},
			Usage:       "calendar URL to use",
			DefaultText: "https://alice.btc.calendar.opentimestamps.org/",
			Value:       cli.NewStringSlice("https://alice.btc.calendar.opentimestamps.org/"),
		},
		&cli.StringSliceFlag{
			Name:    "hash",
			Aliases: []string{"d"},
			Usage:   "32-byte sha256 hash as hex instead of hashing a file (optional)",
		},
	},
	ArgsUsage: "[file]",
	Action: func(c *cli.Context) error {
		var digest [32]byte
		var filename string
		if hash := c.String("hash"); len(hash) == 64 {
			bhash, err := hex.DecodeString(hash)
			if err != nil {
				return fmt.Errorf("invalid 64-char hex string '%s' passed to --hash/-d: %w", hash, err)
			}
			filename = hash
			copy(digest[:], bhash)
		} else if file := c.Args().First(); file != "" {
			b, err := os.ReadFile(file)
			if err != nil {
				return fmt.Errorf("failed to read file '%s': %w", file, err)
			}
			digest = sha256.Sum256(b)
			filename = file
		} else {
			return fmt.Errorf("must either pass a file or a hash digest directly with --hash/-d")
		}

		if dir, err := os.ReadDir(filepath.Dir(filename)); err == nil {
			if slices.ContainsFunc(dir, func(entry os.DirEntry) bool { return entry.Name() == filename+".ots" }) {
				return fmt.Errorf("file '%s.ots' already exists", filename)
			}
		}

		ts := opentimestamps.File{
			Digest:    digest[:],
			Sequences: make([]opentimestamps.Sequence, 0, 5),
		}
		for _, calendarUrl := range c.StringSlice("calendar") {
			seq, err := opentimestamps.Stamp(c.Context, calendarUrl, digest)
			if err != nil {
				fmt.Fprintf(os.Stderr, "- failed to stamp %x at calendar %s: %s", digest, calendarUrl, err)
				continue
			}
			ts.Sequences = append(ts.Sequences, seq)
			fmt.Fprintf(os.Stderr, "- stamped digest %x at calendar %s\n", digest, calendarUrl)
		}

		if len(ts.Sequences) > 0 {
			if err := os.WriteFile(filename+".ots", ts.SerializeToFile(), 0644); err != nil {
				return err
			}
			fmt.Println("saved file '" + filename + ".ots'")
			return nil
		}

		return fmt.Errorf("got not valid stamps")
	},
}

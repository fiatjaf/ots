package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var stamp = &cli.Command{
	Name:        "stamp",
	Usage:       "",
	Description: ``,
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:        "calendar",
			Aliases:     []string{"c"},
			Usage:       "calendar URL to use -- may be specified multiple times",
			DefaultText: "https://alice.btc.calendar.opentimestamps.org/",
		},
		&cli.StringSliceFlag{
			Name:    "hash",
			Aliases: []string{"d"},
			Usage:   "32-byte hash as hex instead of hashing a file (optional)",
		},
	},
	ArgsUsage: "[file]",
	Action: func(c *cli.Context) error {
		fmt.Println(c.StringSlice("calendar"))

		// digest := sha256.Sum256([]byte{1, 2, 3})
		// fmt.Println(hex.EncodeToString(digest[:]))
		// data, err := opentimestamps.Stamp(context.Background(), "https://bob.btc.calendar.opentimestamps.org", digest)
		// if err != nil {
		// 	panic(err)
		// }

		// if err := os.WriteFile("file.ots", data.SerializeToFile(), 0644); err != nil {
		// 	panic(err)
		// }

		return nil
	},
}

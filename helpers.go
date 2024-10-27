package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v3"
)

func getStdin() string {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		read := bytes.NewBuffer(make([]byte, 0, 1000))
		_, err := io.Copy(read, os.Stdin)
		if err == nil {
			return read.String()
		}
	}
	return ""
}

func getStdinBytes() []byte {
	stdin := getStdin()
	if b, err := hex.DecodeString(stdin); err == nil {
		return b
	}
	if b, err := base64.StdEncoding.DecodeString(stdin); err == nil {
		return b
	}
	return []byte(stdin)
}

func readFromStdinOrFile(c *cli.Command) ([]byte, error) {
	b := getStdinBytes()
	if len(b) > 0 {
		return b, nil
	}

	file := c.Args().First()
	if file == "" {
		return nil, fmt.Errorf("no files or stdin given")
	}

	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return b, nil
}

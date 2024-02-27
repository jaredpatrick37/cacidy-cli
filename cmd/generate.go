package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/jaredpatrick37/cacidy-cli/pkg/generator"
	"github.com/urfave/cli/v3"
)

var generateCommand = &cli.Command{
	Name:  "generate",
	Usage: "generate a new cacidy project",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "sdk",
			Usage:    fmt.Sprintf("pipeline sdk (%s)", strings.Join(generator.SDKS, ", ")),
			Required: true,
		},
		&cli.StringFlag{
			Name:  "go-version",
			Usage: "golang version",
			Value: generator.DefaultGoVersion,
		},
	},
	ArgsUsage: "[path]",
	Action: func(_ context.Context, c *cli.Command) error {
		return generator.New(c.Args().First(), c.String("sdk"), generator.GenerateArgs{
			GoVersion: c.String("go-version"),
		})
	},
}

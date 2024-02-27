package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/jaredpatrick37/cacidy-cli/pkg/project"
	"github.com/urfave/cli/v3"
)

var generateCommand = &cli.Command{
	Name:  "generate",
	Usage: "generate a new cacidy project",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "sdk",
			Usage:    fmt.Sprintf("pipeline sdk (%s)", strings.Join(project.SDKS, ", ")),
			Required: true,
		},
		&cli.StringFlag{
			Name:  "go-version",
			Usage: "golang version",
			Value: project.DefaultGoVersion,
		},
	},
	ArgsUsage: "[path]",
	Action: func(_ context.Context, c *cli.Command) error {
		return project.New(c.Args().First(), c.String("sdk"), project.NewProjectArgs{
			GoVersion: c.String("go-version"),
		})
	},
}

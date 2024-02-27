package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

var rootCmd = &cli.Command{
	Name: "cacidy",
	Commands: []*cli.Command{
		generateCommand,
	},
}

func main() {
	if err := rootCmd.Run(context.Background(), os.Args); err != nil {
		fmt.Printf("%s %s\n", errorStyle.Render("[ERR]"), err)
	}
}

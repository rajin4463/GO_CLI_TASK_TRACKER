package main

import (
	"log"
	"os"
    "context"
    "github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name: "list",
		Usage: "List all tasks",
		Action: func(ctx context.Context, c *cli.Command) error {
            data := csvread("tasks.csv")
            formatTab(data)
            return nil
        },
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}
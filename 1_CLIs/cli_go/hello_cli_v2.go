package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:  "hello_cli",
		Usage: "Print Hello World",
		Action: func(c *cli.Context) error {
			fmt.Printf("Hello")
			return nil
		},
	}

	app.Run(os.Args)
}

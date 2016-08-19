package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var language string

	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang",
			Value:       "english",
			Usage:       "language for the greeting",
			Destination: &language,
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Printf("c.NArg() = %v\n", c.NArg())
		fmt.Printf("c.String(\"lang\") = %v\n", c.String("lang"))
		fmt.Printf("language = %v", language)
		return nil
	}

	app.Run(os.Args)
}

package main

import (
	"log"
	"os"

	"remproject/internal/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "Remproject",
		Action: cmd.Start,
	}

	if err := app.Run(os.Args); err != nil {
		log.Println(err)
	}
}

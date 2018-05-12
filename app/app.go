package app

import (
	"log"
	"os"

	"ToDoist/config"
	"ToDoist/server"
	"github.com/urfave/cli"
)

func init() {
	log.Print("Starting the app!'")
	clientApp := cli.NewApp()

	clientApp.Name = "todoist"
	clientApp.Version = "1.0.0"

	clientApp.Commands = []cli.Command{
		{
			Name:        "run",
			Description: "Starts the server",
			Action: func(c *cli.Context) {
				server.StartServer(config.Port())
			},
		},
	}

	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}

}

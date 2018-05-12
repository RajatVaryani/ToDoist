package app

import (
	"github.com/urfave/cli"
	"os"
	"log"
	"github.com/ToDoist/server"
	"github.com/ToDoist/config"
)

func init()  {
	log.Print("Starting the app!'")
	clientApp := cli.NewApp()

	clientApp.Name = "todoist"
	clientApp.Version = "1.0.0"

	clientApp.Commands = []cli.Command{
		{
			Name: "run",
			Description:"Starts the server",
			Action: func(c *cli.Context) {
				server.StartServer(config.Port())
			},
		},
	}

	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}

}

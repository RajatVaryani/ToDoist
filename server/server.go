package server

import (
	"ToDoist/router"
	"github.com/urfave/negroni"
	"github.com/gorilla/context"
	"strconv"
	"os"
	"os/signal"
	"syscall"
	"github.com/tylerb/graceful"
	"time"
)

const timeout = time.Second * 5

const exitCode = 0

func StartServer(port int) {
	todoistRouter := router.Router()
	server := negroni.New()
	server.UseHandler(context.ClearHandler(todoistRouter))
	portInfo := ":" + strconv.Itoa(port)
	signalChannel := make(chan os.Signal)

	signal.Notify(signalChannel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		graceful.Run(portInfo, timeout, todoistRouter)
	}()

	for {
		sig := <-signalChannel
		switch sig {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
			os.Exit(exitCode)
		}
	}
}

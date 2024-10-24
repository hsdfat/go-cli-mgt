package main

import (
	"fmt"
	"go-cli-mgt/pkg/server"
	"go-cli-mgt/pkg/service/cronjob"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	httpServer := server.Initialize()

	go server.ListenAndServe(httpServer)
	go cronjob.InitCronjobList()
	go server.StartTCPServer()
	stopOrKillServer()
}

func stopOrKillServer() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT, os.Interrupt)
	sig := <-signals
	fmt.Println("Receive Signal from OS - Release resource")
	fmt.Println(sig)
	os.Exit(1)
}

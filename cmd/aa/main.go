package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/hsdfat/go-cli-mgt/pkg/cronjob"
	"github.com/hsdfat/go-cli-mgt/pkg/server"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config.env", "config file name, default to config.env")
	flag.Parse()
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFile = filepath.Join(pwd, configFile)

	httpServer := server.Initialize(configFile)

	go server.ListenAndServe(httpServer)
	go cronjob.InitCronjobList()
	go server.StartTCPServer()
	stopOrKillServer()
}

func stopOrKillServer() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	sig := <-signals
	fmt.Println("Receive Signal from OS - Release resource")
	fmt.Println(sig)
	os.Exit(1)
}

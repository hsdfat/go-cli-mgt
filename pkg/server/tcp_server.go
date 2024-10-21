package server

import (
	"bufio"
	"go-cli-mgt/pkg/config"
	"go-cli-mgt/pkg/logger"
	"net"
	"os"
)

func StartTCPServer() {
	tcpPort := ":" + config.GetServerConfig().TcpPort
	listener, err := net.Listen("tcp", tcpPort)
	if err != nil {
		logger.Logger.Error("Error starting TCP server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	logger.Logger.Info("TCP server listening on port ", tcpPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Logger.Error("Error accepting connection:", err)
			continue
		}

		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	logger.Logger.Info("Client connected:", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			logger.Logger.Error("Error reading from client:", err)
			return
		}

		logger.Logger.Infof("Message from %s: %s", conn.RemoteAddr().String(), message)

		_, err = conn.Write([]byte("Message received: " + message))
		if err != nil {
			logger.Logger.Error("Error writing to client:", err)
			return
		}
	}
}

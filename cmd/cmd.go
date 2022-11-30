package cmd

import (
	"mdns/internal/resolver"
	"mdns/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	ress := resolver.New([]string{"1.1.1.1:53", "8.8.8.8:53"})
	//fmt.Println(ress)
	server := server.New(":1488", ress)
	server.Run()

	// Wait for SIGINT or SIGTERM
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	//udpServer.Shutdown()
	//tcpServer.Shutdown()
}

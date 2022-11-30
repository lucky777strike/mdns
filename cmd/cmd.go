package cmd

import (
	"log"
	"mdns/internal/resolver"
	"mdns/internal/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
)

func Run() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if viper.GetStringSlice("servers") == nil {
		log.Fatal("no servers in config file")
	}

	if viper.GetStringSlice("logport") == nil {
		log.Fatal("no servers in config file")
	}

	ress := resolver.New(viper.GetStringSlice("servers"), viper.GetString("logport"))

	if viper.GetString("servaddr") == "" {
		log.Fatal("bad servaddr in config file")
	}

	server := server.New(viper.GetString("servaddr"), ress)
	server.Run()

	// Wait for SIGINT or SIGTERM
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	//udpServer.Shutdown()
	//tcpServer.Shutdown()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

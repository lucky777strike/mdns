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

	viper.SetDefault("servers", []string{"8.8.8.8:53"})
	viper.SetDefault("servaddr", ":1234")
	viper.SetDefault("logport", ":8080")

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

package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"test/cmd/config"
	"test/cmd/routers"
	"test/pkg/database"
	"test/pkg/server"
)

func Run() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	cfg := config.LoadConfig()
	log.Println("config loaded")

	dbClient, err := database.InitMongoDB(cfg.DBConfig)
	if err != nil {
		log.Println(err)
	}
	defer dbClient.Disconnect(context.Background())
	db := dbClient.Database(cfg.DBConfig.Name)
	log.Println("success initialize database")

	server := server.New(&cfg.ServerConfig)
	router := server.Router()

	routers.RegisterRouter(router, db)

	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	serveChan := server.Run()
	go server.Stop()

	select {
	case err := <-serveChan:
		if err != nil {
			panic(err)
		}
	case <-sigChan:
	}
}

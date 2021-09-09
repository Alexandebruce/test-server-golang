package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	postgresConnString = "postgres://postgres:8888888@192.168.0.10/calc?sslmode=disable&search_path=test_scheme"
	mongoConnString = "mongodb://127.0.0.1:27017"
)

func main() {
	log.SetFlags(log.Flags() | log.Llongfile)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	app := createApp(postgresConnString, mongoConnString)
	app.Run()

	<-stop
	GracefullyShutdownApp(app)
}

func GracefullyShutdownApp(application *Application) {
	//application.Stop()
}
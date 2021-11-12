package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/henprasetya/news/pkg/lib/repo"
	"github.com/henprasetya/news/pkg/server"
	service "github.com/henprasetya/news/pkg/services"
)

func main() {

	repo := repo.CreateRepository()
	service := service.CreateService(repo)
	httpServer := server.NewHttpServer(service, 8080)

	go httpServer.Run()

	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	select {
	case o := <-term:
		log.Printf("exiting gracefully %s", o.String())
	case er := <-httpServer.ListenError():
		log.Printf("error starting http server, exiting gracefully %s", er.Error())
	}
}

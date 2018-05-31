package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/iceoskara/goworkshop/pkg/routing"
	"github.com/iceoskara/goworkshop/webserver"
)

// go run ./cmd/gophercon/gophercon.go
// curl -i http://127.0.0.1:8000/home
func main() {
	log.Printf("Service is starting...")

	shutdown := make(chan error, 2)
	//
	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("Service port was not set ")

	}
	r := routing.BaseRouter()

	ws := webserver.New("", port, r)

	go func() {
		err := ws.Start()
		shutdown <- err
	}()

	internalport := os.Getenv("INTERNAL_PORT")
	if len(internalport) == 0 {
		log.Fatal("iNTERNAL Service port was not set ")
	}
	diagrouter := routing.DiagnosticsRouter()
	diagserver := webserver.New("", internalport, diagrouter)

	go func() {
		err := diagserver.Start()
		shutdown <- err
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	//killSignal := <- interrupt
	//log.Printf("Got %s . Stopping ..", killSignal)

	select {
	case killSignal := <-interrupt:
		log.Printf("Got %s . Stopping ..", killSignal)
	case err := <-shutdown:
		log.Printf("Got an  error %s . Stopping ..", err)
	}
	// stop servers and task
	log.Print(ws.Stop())
	log.Print(diagserver.Stop())

	// stop extra tasks ...
	log.Print("Service was stopped!")
}

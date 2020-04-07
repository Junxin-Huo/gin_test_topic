package src

import (
	"log"
	"os"
	"os/signal"
)

var ServerSigChan chan os.Signal

func init() {
	ServerSigChan = make(chan os.Signal)
}

func ShutDownServer(err error) {
	log.Print(err)
	ServerSigChan <- os.Interrupt
}

func ServiceNotify() {
	signal.Notify(ServerSigChan, os.Interrupt)
	<- ServerSigChan
}

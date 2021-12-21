package main

import (
	"os"
	"os/signal"
)

func init() {

}

func main() {

	//wait exit
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}

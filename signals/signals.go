package signals

import (
	"log"
	"os"
	"os/signal"
)

type CleanupFunctionType func()

var c chan os.Signal
var cleanUps []CleanupFunctionType

func init() {
	startSignalHandler()
}

/* Start signal manager */
func startSignalHandler() {
	c = make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	log.Println("Signal handler started")
	go handleSignals()
}

/* Add a function to the cleanup queue process */
func AddCleanupFunction(functionName CleanupFunctionType) {
	cleanUps = append(cleanUps, functionName)
	return
}

/* Manage OS signals */
func handleSignals() {
	select {
	case sig := <-c:

		log.Printf("Got %s signal. Aborting...\n", sig)

		/* Run cleanups */
		for _, task := range cleanUps {
			task()
		}

		/* Exit */
		os.Exit(1)
	}
}

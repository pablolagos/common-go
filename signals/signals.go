package signals

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type SignalHandlerFunction func()

type handlers struct {
	signal  os.Signal
	handler SignalHandlerFunction
}

var c chan os.Signal
var signalHandlers []handlers

func init() {
	startSignalHandler()
}

/* Iniciar el manejador de señales */
func startSignalHandler() {
	c = make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, os.Interrupt, os.Kill)
	log.Println("Signal handler started")
	go handleSignals()
}

// Add function to interrupt signal SIGTERM
func AddCleanupFunction(functionName SignalHandlerFunction) {
	AddSignalHandler(os.Interrupt, functionName)
}

// Add a handler attached to specific signal:
// hangup -> HUP
func AddSignalHandler(signal os.Signal, functionName SignalHandlerFunction) {
	handler := handlers{
		signal:  signal,
		handler: functionName,
	}
	signalHandlers = append(signalHandlers, handler)
}

/* Maneja las señales en segundo plano */
func handleSignals() {
	for {
		sig := <-c

		log.Printf("Got '%v' signal...\n", sig.String())

		/* Ejecutar handlers */
		for _, task := range signalHandlers {
			if task.signal == sig {
				task.handler()
			}
		}

		/* Salir */
		if sig == os.Interrupt || sig == os.Kill {
			log.Println("Interrupt signal recieved")
			os.Exit(1)
		}
	}
}

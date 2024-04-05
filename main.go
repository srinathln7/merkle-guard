package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/srinathln7/merkle_gaurd/cmd"
	"github.com/srinathln7/merkle_gaurd/internal/server"
)

func init() {
	cmd.SetupFlags()
}

func main() {

	var runServerInBackground = flag.Bool("server", false, "run grpc server in the background")
	flag.Parse()

	// Check if the --server flag is set
	if *runServerInBackground {
		// Create and start the gRPC server in the background
		// To do this, we spin up a new go routine
		go server.RunServer()

		// Wait for a graceful shutdown signal (e.g., Ctrl+C)
		// This will keep the main function running and prevent it from exiting immediately
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		// If the server is running, return to prevent executing Cobra commands
		return
	}

	//  Execute the Cobra commands otherwise
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal("error:", err)
	}
}

# Package main

The `main.go` file serves as the entry point for the application and is responsible for coordinating the execution of different components based on command-line arguments. Let's dissect its functionality:

1. **Initialization**:
   - It initializes command-line flags using the `flag` package and sets up flags defined in `cmd.SetupFlags()`.

2. **Main Function**:
   - Parses command-line flags, particularly the `--server` flag, which determines whether to run the gRPC server in the background or execute Cobra commands.

   - If the `--server` flag is set (`*runServerInBackground` is `true`):
     - It starts the gRPC server in the background by spinning up a new goroutine that executes `server.RunServer()`.
     - Listens for a graceful shutdown signal (e.g., Ctrl+C) using a channel (`c`).
     - Waits for the signal to be received (`<-c`), effectively blocking the main goroutine from exiting immediately.
     - Upon receiving the signal, the main goroutine returns, allowing the program to exit gracefully.

   - If the `--server` flag is not set (default behavior):
     - It executes Cobra commands defined in `cmd.RootCmd` using `cmd.RootCmd.Execute()`.
     - If an error occurs during command execution, it logs the error and exits the program.

This file orchestrates the behavior of the application, allowing it to function either as a standalone gRPC server or as a CLI tool to execute various commands defined by Cobra. It provides flexibility in how the application is utilized, catering to both server-side and client-side functionalities.
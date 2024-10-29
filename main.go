package main

import (
	"fmt"
	"os"

	"github.com/your_project/cli"
	"github.com/your_project/utils"
)

// main is the entry point of the application.
func main() {
	logger := utils.NewLogger()
	logger.LogInfo("Application starting")

	applicationCLI := cli.NewCLI()

	// Run the CLI application
	if err := applicationCLI.Run(); err != nil {
		logger.LogError(fmt.Sprintf("Application encountered an error: %v", err))
		os.Exit(1)
	}

	logger.LogInfo("Application finished successfully")
}

package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/your_project/config"
	"github.com/your_project/server"
	"github.com/your_project/client"
	"github.com/your_project/utils"
)

// CLI is a struct that encapsulates the command-line interface functionalities.
type CLI struct {
	rootCmd *cobra.Command
	config  *config.Config
	logger  *utils.Logger
	server  *server.Server
	client  *client.Client
}

// NewCLI initializes and returns a new CLI instance with default settings.
func NewCLI() *CLI {
	logger := utils.NewLogger()
	cfg := config.NewConfig()
	cli := &CLI{
		rootCmd: &cobra.Command{
			Use:   "app",
			Short: "A simple CLI application",
			Long:  "This application allows you to run a server and client for task management.",
		},
		config: cfg,
		logger: logger,
	}
	cli.setupCommands()
	return cli
}

// setupCommands sets up the CLI commands.
func (cli *CLI) setupCommands() {
	cli.rootCmd.AddCommand(
		&cobra.Command{
			Use:   "start-server",
			Short: "Starts the server",
			RunE: func(cmd *cobra.Command, args []string) error {
				return cli.SetupServer()
			},
		},
		&cobra.Command{
			Use:   "start-client",
			Short: "Starts the client",
			RunE: func(cmd *cobra.Command, args []string) error {
				return cli.SetupClient()
			},
		},
	)
}

// Run executes the CLI application.
func (cli *CLI) Run() {
	if err := cli.config.LoadConfig(); err != nil {
		cli.logger.LogError(fmt.Sprintf("Failed to load configuration: %v", err))
		return
	}
	if err := cli.rootCmd.Execute(); err != nil {
		cli.logger.LogError(fmt.Sprintf("Command execution failed: %v", err))
	}
}

// SetupServer initializes and starts the server.
func (cli *CLI) SetupServer() error {
	serverConfig := cli.config.GetServerConfig()
	cli.server = server.NewServer(serverConfig.Host, serverConfig.Port, cli.logger)
	cli.logger.LogInfo("Server setup complete")
	return cli.server.StartServer()
}

// SetupClient initializes and starts the client.
func (cli *CLI) SetupClient() error {
	clientConfig := cli.config.GetClientConfig()
	cli.client = client.NewClient(clientConfig.ServerAddress, clientConfig.Timeout, cli.logger)
	cli.logger.LogInfo("Client setup complete")
	return cli.client.ConnectToServer()
}

// ManageSessions manages the server-client sessions.
func (cli *CLI) ManageSessions() {
	// Implement session management logic here
	// For now, we just log the session management
	cli.logger.LogInfo("Managing sessions")
}

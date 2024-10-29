package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/your_project/utils"
)

// Client is a struct that encapsulates client functionalities.
type Client struct {
	serverAddress string
	timeout       time.Duration
	logger        *utils.Logger
}

// NewClient initializes and returns a new Client instance with default settings.
func NewClient(serverAddress string, timeout int, logger *utils.Logger) *Client {
	return &Client{
		serverAddress: serverAddress,
		timeout:       time.Duration(timeout) * time.Second,
		logger:        logger,
	}
}

// ConnectToServer establishes a connection to the server.
func (c *Client) ConnectToServer() error {
	c.logger.LogInfo(fmt.Sprintf("Connecting to server at %s", c.serverAddress))
	client := &http.Client{
		Timeout: c.timeout,
	}
	resp, err := client.Get(c.serverAddress)
	if err != nil {
		c.logger.LogError(fmt.Sprintf("Failed to connect to server: %v", err))
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("Server returned non-OK status: %s", resp.Status)
		c.logger.LogError(errMsg)
		return fmt.Errorf(errMsg)
	}

	c.logger.LogInfo("Successfully connected to server")
	return nil
}

// ExecuteTask executes a task based on the taskID.
func (c *Client) ExecuteTask(taskID string) error {
	c.logger.LogInfo(fmt.Sprintf("Executing task with ID: %s", taskID))
	// Implement task execution logic here
	// For now, we just log the task execution
	return nil
}

// SendData sends data to the server.
func (c *Client) SendData(data string) error {
	c.logger.LogInfo(fmt.Sprintf("Sending data to server: %s", data))
	client := &http.Client{
		Timeout: c.timeout,
	}
	resp, err := client.Post(c.serverAddress, "application/json", nil) // Replace nil with actual data payload
	if err != nil {
		c.logger.LogError(fmt.Sprintf("Failed to send data: %v", err))
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("Server returned non-OK status: %s", resp.Status)
		c.logger.LogError(errMsg)
		return fmt.Errorf(errMsg)
	}

	c.logger.LogInfo("Data sent successfully")
	return nil
}

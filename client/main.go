package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"

	"remote-shell-rpc/shared"
)

func main() {
	// Get server address from command line or use default
	serverAddr := "localhost:8080"
	if len(os.Args) > 1 {
		serverAddr = os.Args[1]
	}

	// Connect to RPC server
	fmt.Printf("🔌 Connecting to server at %s...\n", serverAddr)
	client, err := rpc.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer client.Close()

	fmt.Println("Connected to server successfully!")
	fmt.Println()
	fmt.Println("")
	fmt.Println("      Remote Shell RPC Client          ")
	fmt.Println("")
	fmt.Println("Type commands to execute on remote server")
	fmt.Println("Type 'exit' or 'quit' to disconnect")
	fmt.Println("")
	fmt.Println()

	// Interactive command loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Display prompt
		fmt.Print("remote-shell> ")

		// Read user input
		if !scanner.Scan() {
			break
		}

		command := strings.TrimSpace(scanner.Text())

		// Check for exit commands
		if command == "exit" || command == "quit" {
			fmt.Println(" Disconnecting from server...")
			break
		}

		// Skip empty commands
		if command == "" {
			continue
		}

		// Prepare RPC request
		req := &shared.CommandRequest{Command: command}
		res := &shared.CommandResponse{}

		// Call RPC method
		err := client.Call(shared.MethodName, req, res)
		if err != nil {
			fmt.Printf(" RPC Error: %v\n\n", err)
			continue
		}

		// Display results
		if res.Error != "" {
			fmt.Printf(" Execution Error: %s\n", res.Error)
		}

		if res.Stdout != "" {
			fmt.Print(res.Stdout)
		}

		if res.Stderr != "" {
			fmt.Printf("  Error Output:\n%s", res.Stderr)
		}

		fmt.Printf("Exit Code: %d\n", res.ExitCode)
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading input:", err)
	}
}

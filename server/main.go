package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os/exec"
	"runtime"

	"remote-shell-rpc/shared"
)

// ShellService provides RPC methods for remote shell execution
type ShellService struct{}

// ExecuteCommand executes a shell command and returns the output
func (s *ShellService) ExecuteCommand(req *shared.CommandRequest, res *shared.CommandResponse) error {
	log.Printf("Executing command: %s", req.Command)

	// Determine shell based on OS
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", req.Command)
	} else {
		cmd = exec.Command("sh", "-c", req.Command)
	}

	// Capture stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Execute the command
	err := cmd.Run()

	// Populate response
	res.Stdout = stdout.String()
	res.Stderr = stderr.String()

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			res.ExitCode = exitErr.ExitCode()
		} else {
			res.ExitCode = -1
			res.Error = err.Error()
		}
		log.Printf("Command failed: %v", err)
	} else {
		res.ExitCode = 0
		log.Printf("Command executed successfully")
	}

	return nil // Always return nil to RPC framework (errors in res.Error)
}

func main() {
	// Create and register the RPC service
	shellService := new(ShellService)
	err := rpc.Register(shellService)
	if err != nil {
		log.Fatal("Error registering RPC service:", err)
	}

	// Listen on TCP port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error starting TCP listener:", err)
	}
	defer listener.Close()

	fmt.Println("")
	fmt.Println("   Remote Shell RPC Server Started     ")
	fmt.Println("")
	fmt.Println(" Listening on port 8080...")
	fmt.Println(" Waiting for client connections...")
	fmt.Println()

	// Accept and handle client connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		log.Printf("New client connected: %s", conn.RemoteAddr())

		// Handle each client in a separate goroutine
		go rpc.ServeConn(conn)
	}
}

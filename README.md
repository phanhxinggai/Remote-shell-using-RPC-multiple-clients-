# Remote Shell using RPC (Multiple Clients)

A distributed remote shell system built with Golang's native RPC framework that allows multiple clients to connect to a central server and execute shell commands remotely.

## 🎯 Overview

This project demonstrates key distributed systems concepts:
- **RPC (Remote Procedure Call)**: Client-server communication using Go's `net/rpc`
- **Concurrent Client Handling**: Server handles multiple clients simultaneously using goroutines
- **Remote Command Execution**: Execute shell commands on remote server safely
- **Cross-Platform Support**: Works on Windows and Linux/Unix systems

## 🏗️ Architecture

```
┌─────────────┐         RPC          ┌─────────────┐
│  Client 1   │◄────────────────────►│             │
└─────────────┘                      │             │
                                     │             │
┌─────────────┐         RPC          │   Server    │
│  Client 2   │◄────────────────────►│             │
└─────────────┘                      │             │
                                     │             │
┌─────────────┐         RPC          │             │
│  Client N   │◄────────────────────►│             │
└─────────────┘                      └─────────────┘
                                           │
                                           ▼
                                     ┌─────────────┐
                                     │    Shell    │
                                     │  (OS/Exec)  │
                                     └─────────────┘
```

## 📁 Project Structure

```
remote-shell-rpc/
├── server/
│   └── main.go          # RPC server implementation
├── client/
│   └── main.go          # RPC client implementation
├── shared/
│   └── types.go         # Shared RPC types and interfaces
├── go.mod               # Go module file
└── README.md            # This file
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- Windows, Linux, or macOS

### Running the Server

1. Open a terminal in the project directory
2. Start the server:

```bash
cd server
go run main.go
```

The server will start listening on port 8080.

### Running the Client

1. Open a new terminal (you can open multiple terminals for multiple clients)
2. Start a client:

```bash
cd client
go run main.go
```

Or connect to a remote server:

```bash
cd client
go run main.go <server-ip>:8080
```

### Example Usage

After connecting, you can execute commands:

```
remote-shell> pwd
/home/user/projects
Exit Code: 0

remote-shell> echo "Hello from RPC!"
Hello from RPC!
Exit Code: 0

remote-shell> ls -la
total 20
drwxr-xr-x 5 user user 4096 Dec  6 14:00 .
drwxr-xr-x 3 user user 4096 Dec  6 13:55 ..
...
Exit Code: 0

remote-shell> exit
👋 Disconnecting from server...
```

## 🧪 Testing Multiple Clients

To test concurrent client connections:

1. Start the server in one terminal
2. Open 3-5 additional terminals
3. Start a client in each terminal
4. Execute commands from different clients simultaneously
5. Observe that all clients receive correct responses

## 🔧 Technical Details

### RPC Communication

- **Protocol**: TCP
- **Port**: 8080
- **Service**: `ShellService`
- **Method**: `ExecuteCommand`

### Request/Response Format

**CommandRequest:**
```go
type CommandRequest struct {
    Command string // Shell command to execute
}
```

**CommandResponse:**
```go
type CommandResponse struct {
    Stdout   string // Standard output
    Stderr   string // Standard error
    ExitCode int    // Exit code (0 = success)
    Error    string // Error message if any
}
```

### Concurrency

The server handles each client connection in a separate goroutine, allowing multiple clients to execute commands concurrently without blocking each other.

### Cross-Platform Shell Execution

- **Windows**: Uses `cmd /C <command>`
- **Linux/Unix**: Uses `sh -c <command>`

## 📝 Notes

- Commands are executed with the same privileges as the server process
- Be careful when running this over public networks (no authentication/encryption)
- For production use, consider adding authentication, encryption, and command whitelisting

## 👥 Authors

Distributed Systems - Group Project

## 📄 License

Educational project for Distributed Systems course.

package shared

// CommandRequest represents a request to execute a shell command
type CommandRequest struct {
	Command string // The shell command to execute
}

// CommandResponse represents the response from executing a shell command
type CommandResponse struct {
	Stdout   string // Standard output from the command
	Stderr   string // Standard error from the command
	ExitCode int    // Exit code of the command (0 = success)
	Error    string // Error message if command execution failed
}

// RPC Service and Method Names
const (
	ServiceName = "ShellService"
	MethodName  = "ShellService.ExecuteCommand"
)

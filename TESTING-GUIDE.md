# Remote Shell RPC - Testing Guide

## 📋 Overview

This guide provides step-by-step instructions to test the Remote Shell RPC system with multiple clients.

**What is this system?**
- A **Remote Shell** that executes commands on a server
- Uses **RPC (Remote Procedure Call)** for communication
- Supports **multiple concurrent clients**

---

## 🔧 Prerequisites

**Install Go:**
1. Download from: https://go.dev/dl/
2. Install for Windows (e.g., `go1.21.x.windows-amd64.msi`)
3. Restart your terminal
4. Verify: `go version`

---

## 🚀 Quick Start

### Step 1: Build the Project

Open terminal in the project folder and run:

```bash
cd "c:\Users\Admin\Desktop\Yearly study material\B3\dis sys\remote-shell-rpc"
go mod tidy
```

### Step 2: Build Server and Client

```bash
# Build server
cd server
go build -o server.exe main.go
cd ..

# Build client
cd client
go build -o client.exe main.go
cd ..
```

**Or use the build script:**
```bash
build.bat
```

---

## 🧪 Test Scenarios

### ✅ Test 1: Single Client Connection

**Purpose:** Verify basic RPC functionality

**Steps:**

**Terminal 1 - Start Server:**
```bash
cd server
go run main.go
```

Expected output:
```
╔════════════════════════════════════════╗
║   Remote Shell RPC Server Started     ║
╚════════════════════════════════════════╝
📡 Listening on port 8080...
🔌 Waiting for client connections...
```

**Terminal 2 - Start Client:**
```bash
cd client
go run main.go
```

Expected output:
```
🔌 Connecting to server at localhost:8080...
✅ Connected to server successfully!

╔════════════════════════════════════════╗
║      Remote Shell RPC Client          ║
╚════════════════════════════════════════╝
Type commands to execute on remote server
Type 'exit' or 'quit' to disconnect
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

remote-shell>
```

**Test Commands:**

```bash
# Test 1.1: Echo command
remote-shell> echo "Hello RPC!"
Hello RPC!
Exit Code: 0

# Test 1.2: Current directory
remote-shell> cd

# Test 1.3: List files (Windows)
remote-shell> dir

# Test 1.4: Show current user
remote-shell> whoami

# Test 1.5: Show date/time
remote-shell> date /t

# Test 1.6: Show IP configuration
remote-shell> ipconfig

# Test 1.7: Create a file
remote-shell> echo test content > test.txt

# Test 1.8: Read the file
remote-shell> type test.txt

# Test 1.9: Delete the file
remote-shell> del test.txt

# Test 1.10: Exit
remote-shell> exit
```

**Expected Result:** ✅ All commands execute successfully and return correct output

---

### ✅ Test 2: Multiple Concurrent Clients

**Purpose:** Verify server handles multiple clients simultaneously via RPC

**Steps:**

**Terminal 1 - Server (already running):**
```bash
cd server
go run main.go
```

**Terminal 2 - Client 1:**
```bash
cd client
go run main.go
```
Execute:
```bash
remote-shell> echo "I am Client 1"
remote-shell> ping localhost -n 3
```

**Terminal 3 - Client 2:**
```bash
cd client
go run main.go
```
Execute:
```bash
remote-shell> echo "I am Client 2"
remote-shell> date /t
```

**Terminal 4 - Client 3:**
```bash
cd client
go run main.go
```
Execute:
```bash
remote-shell> echo "I am Client 3"
remote-shell> whoami
```

**Terminal 5 - Client 4:**
```bash
cd client
go run main.go
```
Execute:
```bash
remote-shell> echo "I am Client 4"
remote-shell> dir
```

**Expected Result:** 
- ✅ Server shows 4 client connections
- ✅ All clients can execute commands simultaneously
- ✅ No client blocks another client
- ✅ Each client gets correct output

**Server should show:**
```
✅ New client connected: 127.0.0.1:xxxxx
✅ New client connected: 127.0.0.1:xxxxx
✅ New client connected: 127.0.0.1:xxxxx
✅ New client connected: 127.0.0.1:xxxxx
```

---

### ✅ Test 3: Error Handling

**Purpose:** Verify RPC error handling for invalid commands

**Test Commands:**

```bash
# Test 3.1: Invalid command
remote-shell> invalidcommand123
⚠️  Error Output:
'invalidcommand123' is not recognized as an internal or external command...
Exit Code: 1

# Test 3.2: Command with syntax error
remote-shell> dir /?/?/?
⚠️  Error Output:
[Error message]
Exit Code: 1

# Test 3.3: Access denied
remote-shell> del C:\Windows\System32\kernel32.dll
⚠️  Error Output:
Access is denied.
Exit Code: 1
```

**Expected Result:** ✅ Server handles errors gracefully, returns error messages via RPC

---

### ✅ Test 4: Network Connection Test

**Purpose:** Test RPC connection to remote server

**Connect to specific IP:**
```bash
cd client
go run main.go 192.168.1.100:8080
```

**Expected Result:** ✅ Client connects to server at specified address via RPC

---

### ✅ Test 5: Long-Running Commands

**Purpose:** Test RPC with commands that take time

```bash
# Test 5.1: Ping (takes ~3 seconds)
remote-shell> ping localhost -n 3

# Test 5.2: Directory scan
remote-shell> dir C:\ /s /b | find /c ":"

# Test 5.3: Wait command
remote-shell> timeout /t 5
```

**Expected Result:** ✅ RPC waits for command completion, returns full output

---

### ✅ Test 6: Special Characters & Quotes

**Purpose:** Test RPC marshaling with special characters

```bash
# Test 6.1: Quotes
remote-shell> echo "Hello World!"

# Test 6.2: Ampersand
remote-shell> echo Test & echo Test2

# Test 6.3: Pipe
remote-shell> dir | find "txt"

# Test 6.4: Redirection
remote-shell> echo test > output.txt & type output.txt

# Test 6.5: Multiple commands
remote-shell> echo First && echo Second
```

**Expected Result:** ✅ All special characters handled correctly by RPC

---

## 📊 Performance Testing

### Load Test: 10 Concurrent Clients

**Script to run 10 clients (PowerShell):**

```powershell
# Start server first
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd server; go run main.go"

# Wait for server to start
Start-Sleep -Seconds 2

# Start 10 clients
for ($i=1; $i -le 10; $i++) {
    Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd client; go run main.go"
    Start-Sleep -Milliseconds 500
}
```

**Test in each client:**
```bash
remote-shell> echo "Client test"
remote-shell> whoami
remote-shell> date /t
```

**Measure:**
- Response time for each command
- Server CPU/Memory usage
- No connection failures

---

## 🎯 Verification Checklist

After testing, verify:

- [ ] ✅ Server starts on port 8080
- [ ] ✅ Client connects to server successfully (RPC connection)
- [ ] ✅ Commands execute and return correct output (RPC call works)
- [ ] ✅ Multiple clients connect simultaneously (concurrent RPC handling)
- [ ] ✅ Error messages are returned properly (RPC error handling)
- [ ] ✅ Exit command disconnects client cleanly (RPC connection close)
- [ ] ✅ Server continues running after client disconnects (RPC server stability)

---

## 🐛 Troubleshooting

### Error: "go: command not found"
**Solution:** Install Go from https://go.dev/dl/ and restart terminal

### Error: "cannot find package"
**Solution:** 
```bash
go mod tidy
```

### Error: "bind: address already in use"
**Solution:** Port 8080 is occupied. Close the existing server or change port in code.

### Error: "connection refused"
**Solution:** Make sure server is running before starting client

### Client freezes
**Solution:** Press Ctrl+C to stop, restart client

---

## 📝 Test Report Template

After testing, document your results:

```
# Test Report

**Date:** [Date]
**Tester:** [Your Name]

## Single Client Test
- Status: PASS / FAIL
- Notes: [Any observations]

## Multiple Clients Test
- Number of clients tested: [X]
- Status: PASS / FAIL
- Notes: [Any observations]

## Error Handling Test
- Status: PASS / FAIL
- Notes: [How errors were handled]

## Performance Observations
- Average response time: [X ms]
- Max concurrent clients: [X]
- Server CPU usage: [X%]
- Server memory usage: [X MB]

## Issues Found
1. [Issue 1]
2. [Issue 2]

## Conclusion
[Overall assessment]
```

---

## 🎓 Understanding the RPC Architecture

**What happens when you type a command:**

1. **Client:** User types command → `CommandRequest` struct created
2. **Client:** Calls RPC method `ShellService.ExecuteCommand(request, response)`
3. **Network:** Request marshaled and sent via TCP to server
4. **Server:** Receives RPC call, unmarshals `CommandRequest`
5. **Server:** Executes shell command using `os/exec`
6. **Server:** Populates `CommandResponse` with stdout, stderr, exit code
7. **Network:** Response marshaled and sent back via TCP
8. **Client:** Receives RPC response, unmarshals `CommandResponse`
9. **Client:** Displays output to user

**This is true RPC because:**
- ✅ Client calls remote procedure (ExecuteCommand) as if it's local
- ✅ Marshaling/Unmarshaling handled automatically by `net/rpc`
- ✅ Network communication abstracted away
- ✅ Server executes and returns result transparently

---

## 📚 Additional Resources

- Go RPC Documentation: https://pkg.go.dev/net/rpc
- Project README: `README.md`
- Setup Guide: `SETUP.md`
- Technical Report: `remote-shell-rpc-report.tex`

---

**Happy Testing! 🚀**

If you encounter any issues, check the troubleshooting section or review the server logs for error messages.

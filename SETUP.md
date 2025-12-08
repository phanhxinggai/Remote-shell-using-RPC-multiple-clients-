# Hướng dẫn Cài đặt và Chạy

## Yêu cầu

Trước khi chạy dự án, bạn cần cài đặt Go (Golang):

1. Tải Go từ: https://go.dev/dl/
2. Chọn phiên bản cho Windows (go1.21.x.windows-amd64.msi)
3. Cài đặt và khởi động lại terminal

Kiểm tra Go đã cài đặt:
```bash
go version
```

## Cách 1: Sử dụng Scripts (Đơn giản nhất)

### Build dự án:
```bash
build.bat
```

### Chạy server (Terminal 1):
```bash
run-server.bat
```

### Chạy client (Terminal 2):
```bash
run-client.bat
```

## Cách 2: Build và chạy thủ công

### Build dự án:
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

### Chạy server (Terminal 1):
```bash
cd server
server.exe
```

### Chạy client (Terminal 2):
```bash
cd client
client.exe
```

## Cách 3: Chạy trực tiếp với go run (không cần build)

### Chạy server (Terminal 1):
```bash
cd server
go run main.go
```

### Chạy client (Terminal 2):
```bash
cd client
go run main.go
```

## Test với nhiều clients

Mở nhiều terminal và chạy client ở mỗi terminal:

**Terminal 1:**
```bash
cd server
go run main.go
```

**Terminal 2, 3, 4, 5:**
```bash
cd client
go run main.go
```

Thử chạy các lệnh trên mỗi client:
- `echo "Hello from client 1"`
- `dir` hoặc `ls`
- `whoami`
- `date`

## Compile LaTeX Report

Để tạo PDF từ file LaTeX:

```bash
pdflatex remote-shell-rpc-report.tex
pdflatex remote-shell-rpc-report.tex
```

Hoặc sử dụng Online LaTeX Editor như Overleaf:
1. Truy cập: https://www.overleaf.com/
2. Tạo project mới
3. Upload file `remote-shell-rpc-report.tex`
4. Compile để tạo PDF

## Troubleshooting

### Lỗi: 'go' is not recognized
- Go chưa được cài đặt hoặc chưa có trong PATH
- Cài đặt Go từ https://go.dev/dl/
- Khởi động lại terminal sau khi cài đặt

### Lỗi: cannot find package "remote-shell-rpc/shared"
```bash
go mod tidy
```

### Server không chạy được
- Kiểm tra port 8080 có bị chiếm không
- Chạy với quyền administrator nếu cần

### Client không kết nối được
- Đảm bảo server đang chạy
- Kiểm tra firewall không block port 8080

## Ví dụ sử dụng

Sau khi connect:

```
remote-shell> echo "Hello RPC!"
Hello RPC!
Exit Code: 0

remote-shell> cd c:\
Exit Code: 0

remote-shell> dir
[Directory listing]
Exit Code: 0

remote-shell> exit
👋 Disconnecting from server...
```

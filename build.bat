@echo off
echo ========================================
echo Building Remote Shell RPC Project
echo ========================================
echo.

echo [1/3] Building Server...
cd server
go build -o server.exe main.go
if %errorlevel% neq 0 (
    echo ERROR: Server build failed!
    exit /b 1
)
echo ✓ Server built successfully
cd ..

echo.
echo [2/3] Building Client...
cd client
go build -o client.exe main.go
if %errorlevel% neq 0 (
    echo ERROR: Client build failed!
    exit /b 1
)
echo ✓ Client built successfully
cd ..

echo.
echo [3/3] Verifying builds...
if exist "server\server.exe" (
    echo ✓ server\server.exe created
) else (
    echo ✗ server\server.exe not found
    exit /b 1
)

if exist "client\client.exe" (
    echo ✓ client\client.exe created
) else (
    echo ✗ client\client.exe not found
    exit /b 1
)

echo.
echo ========================================
echo Build completed successfully!
echo ========================================
echo.
echo To run the server:  cd server ^&^& server.exe
echo To run the client:  cd client ^&^& client.exe
echo.
pause

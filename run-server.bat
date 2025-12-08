@echo off
echo ========================================
echo   Remote Shell RPC Server
echo ========================================
echo Starting server on port 8080...
echo.

cd server
go run main.go

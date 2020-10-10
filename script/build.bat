cd %~dp0
set GOPROXY=https://goproxy.io
set GOOS=windows
set GOARCH=amd64
go build -ldflags "-s -w" -o ../bin/clagent_windows_amd64.exe ../cmd/main.go
set GOOS=linux
set GOARCH=amd64
go build -ldflags "-s -w" -o ../bin/clagent_linux_amd64 ../cmd/main.go
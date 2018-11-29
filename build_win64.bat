@echo off
@color 06

set GoDevWork=%cd%\

set GOOS=windows
set GOARCH=amd64
set GOPATH=%GoDevWork%;%GOPATH%

echo "Build For ha666_server ..."

cd %GoDevWork%ha666_server\
go build -ldflags "-s -w" -o ha666_server.exe

echo "--------- Build For ha666_server Success!"

echo "Build For ha666_gateway ..."

cd %GoDevWork%ha666_gateway\
go build -ldflags "-s -w"

echo "--------- Build For ha666_gateway Success!"

echo "Build For ha666_client ..."

cd %GoDevWork%ha666_client\
go build -ldflags "-s -w" -o ha666_client.exe

echo "--------- Build For ha666_client Success!"

echo "Build For ha666web ..."

cd %GoDevWork%ha666web\
go build -ldflags "-s -w" -o ha666web.exe

echo "--------- Build For ha666web Success!"

pause
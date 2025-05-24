@echo off

set GOPATH=^
%USERPROFILE%\Downloads\gopath;

set GOROOT=^
%USERPROFILE%\Downloads\go1.23.0.windows-amd64\go

set PATH=^
%GOROOT%\bin;

echo %PATH%

set GOOS=windows
set GOARCH=amd64
go get github.com/yeka/zip
go build

@REM SET GOOS=js
@REM SET GOARCH=wasm

@REM go build -o build-webassembly\main.wasm  .\webassembly\main_js.go

@REM set GOROOT=%USERPROFILE%\Downloads\go
@REM copy "%GOROOT%\misc\wasm\wasm_exec.js" build-webassembly\
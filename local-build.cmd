@echo off

set GOPATH=^
%USERPROFILE%\Downloads\gopath;

set GOROOT=^
%USERPROFILE%\Downloads\go1.25.0.windows-amd64\go

set PATH=^
%GOROOT%\bin;

echo %PATH%

set GOOS=windows
set GOARCH=amd64
go build

SET GOOS=js
SET GOARCH=wasm

go build -o build-webassembly\main.wasm  .\webassembly\main_js.go

copy "%GOROOT%\lib\wasm\wasm_exec.js" build-webassembly\
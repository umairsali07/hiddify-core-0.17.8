@echo off
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=1
del bin\libcore.dll bin\BarVPNCli.exe
set CGO_LDFLAGS=
go build -trimpath -tags with_gvisor,with_quic,with_wireguard,with_ech,with_utls,with_clash_api,with_grpc -ldflags="-w -s" -buildmode=c-shared -o bin/libcore.dll ./custom
go get github.com/akavel/rsrc
go install github.com/akavel/rsrc

rsrc  -ico .\assets\hiddify-cli.ico -o cli\bydll\cli.syso

copy bin\libcore.dll .
set CGO_LDFLAGS="libcore.dll"
go build  -o bin/BarVPNCli.exe ./cli/bydll/
del libcore.dll


:: Wait for a key press
pause
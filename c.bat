del nm.exe
go build -o nm_cmd.exe ip.go systray.go conf.go pic.go
go build -ldflags -H=windowsgui -o nm.exe ip.go systray.go conf.go pic.go
::go build -o tcp_seq_check.exe
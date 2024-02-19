del nm.exe
::go build -o nm-cmd.exe ip.go systray.go conf.go pic.go win.go
::go build -ldflags "-w -s -H windowsgui" -o nm.exe ip.go systray.go conf.go pic.go win.go
go build -ldflags "-w -s -H windowsgui" -o nm.exe
go build -o nm-cmd.exe
::go build -o tcp_seq_check.exe
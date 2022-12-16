package main

import (
	"awesomeProject/internal/tea"
	"awesomeProject/internal/udp"
)

var cmdLine = make(chan []byte)
var logLine = make(chan []byte)

func main() {
	//app.Words()
	//utils.Test()
	//network.Tcp()
	//network.Scan()

	go udp.InitUdp(cmdLine, logLine)
	tea.Start(cmdLine, logLine)
}

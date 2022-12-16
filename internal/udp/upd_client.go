package udp

import (
	"awesomeProject/internal/utils"
	"bytes"
	"fmt"
	"log"
	"net"
)

// UdpClient for drone
type UdpClient struct {
	c           *net.UDPConn
	response    chan []byte
	logger      *log.Logger
	lastMassage []byte
}

func NewUdpClient(adr net.UDPAddr) *UdpClient {
	target := UdpClient{
		response: make(chan []byte),
		logger:   utils.InitLogger("UDP logger"),
	}

	// Create connection
	c, err := net.DialUDP("udp", nil, &adr)
	if err != nil {
		target.logger.Println("Error when create UdpClient:", err)
		return nil
	}
	target.c = c

	return &target
}

func (u *UdpClient) sendData(data []byte) {
	_, err := u.c.Write(data)
	if err != nil {
		u.logger.Println("Sending error:", err)
	}
}

func (u *UdpClient) readUdp() []byte {
	b := make([]byte, 100)
	oop := make([]byte, 100)

	n, _, _, _, err := u.c.ReadMsgUDP(b, oop)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	b = b[0:n]
	if bytes.Equal(u.lastMassage, b) {
		u.lastMassage = b
		u.logger.Println("Last new :", b)
	}

	return b
}

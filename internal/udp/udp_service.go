package udp

import (
	"time"
)

var b4 = []byte{0xef, 0x00, 0x04, 0x00}

// UdpService upd service
type UdpService struct {
	client *UdpClient
	cl168  *UdpClient

	cmdLine chan []byte
	logLine chan []byte
}

func newUdpService(client *UdpClient, cl168 *UdpClient, cmdLine chan []byte, logLine chan []byte) *UdpService {
	return &UdpService{client: client, cl168: cl168, cmdLine: cmdLine, logLine: logLine}
}

func InitUdp(cmdLine chan []byte, logLine chan []byte) {

	s := newUdpService(NewUdpClient(droneUDPAdr), NewUdpClient(droneUDPAdr0), cmdLine, logLine)

	go func() {
		for {
			s.cl168.sendData(b4)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		err := s.listener()
		if err != nil {
			s.client.logger.Println("Listener error:", err)
			return
		}
	}()

	for {
		if cmd := <-cmdLine; true {
			s.logLine <- cmd
			s.client.sendData(cmd)
		}
	}
}

func (s *UdpService) listener() error {
	for {
		response := s.client.readUdp()
		if len(response) != 0 {
			s.logLine <- response
		}
	}
}

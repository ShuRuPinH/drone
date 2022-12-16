package network

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var droneAdr = net.TCPAddr{
	IP:   net.IP{192, 168, 1, 1},
	Port: 7070,
	Zone: "",
}

var stubAdr = net.TCPAddr{
	IP:   net.IP{45, 79, 112, 203},
	Port: 4242,
	Zone: "",
}

func Tcp() {

	args := os.Args
	for i, arg := range args {
		fmt.Println(i, arg)
	}

	c, err := net.DialTCP("tcp", nil, &droneAdr)
	if err != nil {
		fmt.Println(err)
		return
	}

	opt := "OPTIONS rtsp://192.168.1.1:7070/webcam RTSP/1.0\r\nCSeq: 1\r\nUser-Agent: Lavf57.71.100\r\n\r\n"
	des := "DESCRIBE rtsp://192.168.1.1:7070/webcam RTSP/1.0\r\nAccept: application/sdp\r\nCSeq: 2\r\nUser-Agent: Lavf57.71.100\r\n\r\n"
	set := "SETUP rtsp://192.168.1.1:7070/webcam/track0 RTSP/1.0\r\nTransport: RTP/AVP/UDP;unicast;client_port=18798-18799\r\nCSeq: 3\r\nUser-Agent: Lavf57.71.100\r\n\r\n"

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		switch strings.TrimSpace(string(text)) {
		case "11":
			text = opt
		case "22":
			text = des
		case "33":
			text = set

		}

		//_, err = c.Read([]byte(text + "\n"))
		//fmt.Fprintf(c, text+"\n")
		fmt.Fprint(c, text)

		message := readAll(c)

		fmt.Print("->: ", string(message))

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}

func readAll(c *net.TCPConn) []byte {
	result := make([]byte, 0)
	reader := bufio.NewReader(c)

	for {
		b, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("readLine err")
		}
		if len(b) == 0 {
			break
		}
		result = append(result, b...)
		fmt.Println(string(b))
	}
	fmt.Println("read done")
	return result
}

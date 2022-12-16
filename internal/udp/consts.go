package udp

import "net"

var udpTest = net.UDPAddr{
	IP:   net.IP{52, 20, 16, 20},
	Port: 40000,
	Zone: "",
}

var droneUDPAdr = net.UDPAddr{
	IP:   net.IP{192, 168, 1, 1},
	Port: 7099,
	Zone: "",
}
var droneUDPAdr0 = net.UDPAddr{
	IP:   net.IP{192, 168, 169, 1},
	Port: 8800,
	Zone: "",
}

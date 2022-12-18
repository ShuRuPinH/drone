package tea

import (
	"strings"
	"time"
)

var (
	stopChan bool
	upChan   = make(chan bool)
	downChan = make(chan bool)
)

var power = 0

func (m *model) messageProcess(cmd string) {
	switch strings.TrimSpace(cmd) {
	case "b4":
		m.cmdLine <- b4
	case "i":
		m.cmdLine <- b2
	case "up":
		m.up()
	case "down":
		m.down()
	case "d":
		m.cmdLine <- def
	case "s":
		m.cmdLine <- b2
		m.cmdLine <- def
		go m.processLoop()
	case "77":
		m.cmdLine <- b2
		m.cmdLine <- def
		for i := 89; i < 233; i += 10 {
			m.cmdLine <- []byte{0x03, 0x66, 0x80, 0x80, byte(i), 0x80, 0x00, byte((i + 128) % 256), 0x99}
			time.Sleep(time.Millisecond * 50)
		}
		for i := 233; i < 30; i -= 10 {
			m.cmdLine <- []byte{0x03, 0x66, 0x80, 0x80, byte(i), 0x80, 0x00, byte((i - 128) % 256), 0x99}
			time.Sleep(time.Millisecond * 50)
		}

	default:
		m.cmdLine <- []byte{0x00, 0x01, 0x02}
	}
}

func (m *model) processLoop() {
	for !stopChan {
		m.cmdLine <- []byte{0x03, 0x66, 0x80, 0x80, byte(power), 0x80, 0x00, byte((power + 128) % 256), 0x99}
		time.Sleep(time.Millisecond * 100)
		//power += 5

		if power < 172 {
			power += 5
		}
	}
}

func (m *model) up() {
	power += 10
}

func (m *model) down() {
	power -= 10
}

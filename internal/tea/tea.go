package tea

// A simple program demonstrating the text area component from the Bubbles
// component library.

import (
	"awesomeProject/internal/utils"
	"bytes"
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"log"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	b2 = []byte{0x01, 0x01}

	def = []byte{0x03, 0x66, 0x80, 0x80, 0x00, 0x80, 0x00, 0x80, 0x99}

	b4 = []byte{0xef, 0x00, 0x04, 0x00}
	//des := "DESCRIBE rtsp://192.168.1.1:7070/webcam RTSP/1.0\r\nAccept: application/sdp\r\nCSeq: 2\r\nUser-Agent: Lavf57.71.100\r\n\r\n"
	//set := "SETUP rtsp://192.168.1.1:7070/webcam/track0 RTSP/1.0\r\nTransport: RTP/AVP/UDP;unicast;client_port=18798-18799\r\nCSeq: 3\r\nUser-Agent: Lavf57.71.100\r\n\r\n"

	t1 = []byte{0x03, 0x66, 0x80, 0x80, 0x59, 0x80, 0x00, 0xd9, 0x99}

	answer = []byte{0x2c, 0x00, 0x00}
)

func Start(cmdLine chan []byte, logLine chan []byte) {
	m := initialModel(cmdLine, logLine)
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

}

type (
	errMsg error
)

type model struct {
	viewport    viewport.Model
	statview    viewport.Model
	helpview    viewport.Model
	messages    []string
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error
	cmdLine     chan []byte
	logLine     chan []byte
	//
	help     help.Model
	quitting bool
	keys     keyMap
}

func initialModel(cmdLine chan []byte, logLine chan []byte) *model {
	ta := textarea.New()
	ta.Placeholder = "Send a command..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetWidth(50)
	ta.SetHeight(1)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(200, 5)
	vp.SetContent(`Logs start here!`)

	sv := viewport.New(30, 3)
	sv.SetContent(`Last drone response...`)

	hv := viewport.New(30, 10)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return &model{
		help:        help.New(),
		keys:        keys,
		statview:    sv,
		viewport:    vp,
		helpview:    hv,
		messages:    []string{},
		textarea:    ta,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#008800")),
		err:         nil,
		cmdLine:     cmdLine,
		logLine:     logLine,
	}

}

func (m *model) Init() tea.Cmd {
	go m.listen()
	return textarea.Blink
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			fmt.Println(m.textarea.Value())
			return m, tea.Quit
		case tea.KeyEnter:
			m.messageProcess(m.textarea.Value())
			m.textarea.Reset()
		//m.viewport.GotoBottom()
		case tea.KeyCtrlY:

			m.addLogMess([]byte(m.textarea.Value()))
			m.textarea.Reset()

		}
		switch {
		case key.Matches(msg, m.keys.Up):
			m.messageProcess("up")
			m.textarea.Reset()

		case key.Matches(msg, m.keys.Down):
			m.messageProcess("down")
			m.textarea.Reset()
		case key.Matches(msg, m.keys.Help):

		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			stopChan = true
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func (m *model) addLogMess(data []byte) {
	tmpMess := m.senderStyle.Render(time.Now().Format(time.StampMilli)+" ") + utils.HexDataString(data)
	if len(m.messages) == m.viewport.Height {
		m.messages = m.messages[:m.viewport.Height-1]
	}
	m.messages = append([]string{tmpMess}, m.messages...)
	//	fmt.Println(m.messages)
	m.viewport.SetContent(strings.Join(m.messages, "\n"))
}

func (m *model) View() string {
	///
	if m.quitting {
		return "Bye!\n"
	}
	///

	// MainView
	mainView := lipgloss.JoinHorizontal(0, m.textarea.View(),
		m.viewport.View(), "\t\t\t\t", m.statview.View()) + "\n"

	return mainView + m.help.View(m.keys)

}

func (m *model) listen() {
	for {
		if logText := <-m.logLine; true {
			if bytes.Equal(logText, answer) {

				m.statview.SetContent(lipgloss.NewStyle().
					Foreground(lipgloss.Color("#d7b93e")).Render(time.Now().Format(time.StampMilli)))
			} else {
				m.addLogMess(logText)
			}
		}
	}
}

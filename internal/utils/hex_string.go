package utils

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

var numStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#e5e23a"))

var dataStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#197288"))

var dataHStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#e56dbd"))

func HexDataString(data []byte) string {
	result := ""

	for i, datum := range data {
		tmpNum := numStyle.Render(fmt.Sprintf("%d", i))
		tmpData := ""

		if i == 4 || i == 7 {
			tmpData = dataHStyle.Render(fmt.Sprintf("[%#x]=%d ", datum, datum))
		} else {
			tmpData = dataStyle.Render(fmt.Sprintf("[%#x] ", datum))
		}

		result += tmpNum + tmpData
	}
	return result
}

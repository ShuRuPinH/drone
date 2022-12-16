package utils

import "fmt"

func HexDataString(data []byte) string {
	result := ""
	for i, datum := range data {
		result += fmt.Sprintf("%d[%#x] ", i, datum)
	}
	return result
}

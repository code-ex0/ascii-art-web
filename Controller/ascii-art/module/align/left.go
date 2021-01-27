package align

import (
	"strings"
)

func left(sentence string) (result string) {
	line := strings.Split(sentence, "\n")

	for i := 0; i < len(line)-1; i++ {
		result += line[i]
		if i != len(line)-1 {
			result += "\n"
		}
	}
	return
}

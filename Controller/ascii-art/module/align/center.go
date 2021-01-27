package align

import (
	"strings"
)

func center(sentence string, sizeCmd int) (result string) {
	for _, l := range strings.Split(sentence, " \n") {
		if l != "" {
			result += getSpace((sizeCmd-len(l))/2) + l + "\n"
		}
	}
	return
}

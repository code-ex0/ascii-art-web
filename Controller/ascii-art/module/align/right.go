package align

import (
	"strings"
)

func right(sentence string, sizeCmd int) (result string) {
	for _, l := range strings.Split(sentence, "\n") {
		if l != "" {
			result += getSpace(sizeCmd-len(l)-1) + l + "\n"
		}
	}
	return
}

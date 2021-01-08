package align

import (
	"fmt"
	"strings"
)

func right(sentence string, sizeCmd int) {
	for _, l := range strings.Split(sentence, "\n") {
		if l != "" {
			fmt.Println(getSpace(sizeCmd-len(l)-1), l)
		}
	}
}

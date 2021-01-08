package reverse

import (
	"../struct"
	"fmt"
)

func Reverse(file []string, alphabet *_struct.Alphabet) {
	lastIndex := 0
	file = file[:len(file)-1]
	for lastIndex <= len(file[0])-2 {
		for i := ' '; i < '~'; i++ {
			find := true
			temp := alphabet.LetterAscii[string(i)]
			for j := 0; j < len(file); j++ {
				if len(file[j][lastIndex:]) >= len(temp[j]) {
					if !(file[j][lastIndex:lastIndex+len(temp[j])] == temp[j]) {
						find = false
					}
				} else {
					find = false
				}
			}
			if find {
				lastIndex += len(temp[0])
				fmt.Print(string(i))
				break
			}
		}
	}
	fmt.Print("\n")
}

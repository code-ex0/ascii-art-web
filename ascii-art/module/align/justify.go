package align

import (
	"fmt"
	"strings"
)

func justify(sentence string, sizeCmd int) {
	line := strings.Split(sentence, "\n")
	for i := 0; i < len(line)/8; i++ {
		if line[i] != "" {
			temp := spaceDetection(line[i*8 : (i+1)*8])
			for _, l := range temp {
				if l != "" {
					numbersSpace := strings.Count(l, "\t")
					sizeSentence := len(strings.ReplaceAll(l, "\t", ""))
					space := numberSpace(numbersSpace, sizeCmd, sizeSentence)
					z := 0
					for j, k := range strings.Split(l, "\t") {
						if k != "" {
							if j != 0 {
								fmt.Print(getSpace(space[z]))
								z++
							}
							fmt.Print(k)
						}
					}
					fmt.Println()
				}
			}
		}
	}
}

func spaceDetection(board []string) []string {
	for i := 0; i < len(board[0])-7; i++ {
		temp := true
		for _, l := range board {
			if !index(l[i:i+6], "     ") {
				temp = false
			}
		}
		if temp {
			for j := 0; j < len(board); j++ {
				board[j] = board[j][:i] + "\t" + board[j][i+6:]
			}
			i += 7
		}
	}
	return board
}

func index(s string, toFind string) bool {
	for i := 0; i < len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind {
			return true
		}
	}
	return false
}

func numberSpace(numbersSpace int, sizeCmd int, sizeSentence int) []int {
	var result = make([]int, numbersSpace)
	allSpace := sizeCmd - sizeSentence
	for i := 0; i < numbersSpace; i++ {
		temp := allSpace / numbersSpace
		if allSpace%numbersSpace != 0 {
			result[i] = temp + 1
			allSpace--
		} else {
			result[i] = temp
		}
	}
	return result
}

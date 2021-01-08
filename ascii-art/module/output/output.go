package output

import (
	"log"
	"os"
)

func Output(file string, sentence [8]string) {
	f, err := os.Create("output/" + file)
	if err != nil {
		log.Fatal(err)
	}
	for _, letter := range sentence {
		_, err2 := f.WriteString(letter)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
	f.Close()
}

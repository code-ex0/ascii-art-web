package align

func Align(sentence string, align string, sizeCmd int) (result string) {

	switch align {
	case "left":
		result = left(sentence)
	case "right":
		result = right(sentence, sizeCmd)
	case "center":
		center(sentence, sizeCmd)
	case "justify":
		justify(sentence, sizeCmd)
	default:
		left(sentence)
	}
	return
}

func getSpace(num int) (result string) {
	for j := 0; j < num; j++ {
		result += " "
	}
	return
}

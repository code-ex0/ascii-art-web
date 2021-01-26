package align

func Align(sentence string, align string, sizeCmd int) (result string) {

	switch align {
	case "left":
		result = left(sentence)
	case "right":
		result = right(sentence, sizeCmd)
	case "center":
		result = center(sentence, sizeCmd)
	case "justify":
		result = justify(sentence, sizeCmd)
	default:
		result = left(sentence)
	}
	return
}

func getSpace(num int) (result string) {
	for j := 0; j < num; j++ {
		result += " "
	}
	return
}

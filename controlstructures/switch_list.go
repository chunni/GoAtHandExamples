package main

func main() {
	println("is space:", isSpace(' '))
	println("is space:", isSpace('a'))
}

func isSpace(c byte) bool {
	switch c {
	case ' ', '\n', '\t', '\r':
		return true
	default:
		return false
	}
}

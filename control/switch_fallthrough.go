package main

func main() {
	println("is space:", isSpaceFallThrough(' '))
	println("is space:", isSpaceFallThrough('a'))
}

func isSpaceFallThrough(c byte) bool {
	switch c {
	case ' ':
		println("case  ")
		fallthrough
	case '\n':
		println("case \\n")
		fallthrough
	case '\t':
		println("case \\t")
		fallthrough
	case '\r':
		println("case \\r")
		return true
	default:
		return false
	}
}

package main

func main() {
	c := 'N'

	switch {
	case 'a' <= c && c <= 'z':
		println("lower case letter")
	case 'A' <= c && c <= 'Z':
		println("upper case letter")
	default:
		println("not a letter")
	}
}

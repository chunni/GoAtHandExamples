package main

import "errors"

func marshalNamed(in string) (out string, err error) {
	if len(in) < 2 {
		err = errors.New("invalid argument")
		return
	}

	out = "pseudo output"
	return
}

func main() {
	out, err := marshalNamed("")
	if err != nil {
		println("marshal error:", err.Error())
	} else {
		println(out)
	}

	out, err = marshalNamed("input")
	if err != nil {
		println("marshal error:", err.Error())
	} else {
		println(out)
	}
}

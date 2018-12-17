package main

import "errors"

func marshal(in string) (string, error) {
	if len(in) < 2 {
		return "", errors.New("invalid argument")
	}

	return "pseudo output", nil
}

func main() {
	out, err := marshal("")
	if err != nil {
		println("marshal error:", err.Error())
	} else {
		println(out)
	}

	out, err = marshal("input")
	if err != nil {
		println("marshal error:", err.Error())
	} else {
		println(out)
	}
}

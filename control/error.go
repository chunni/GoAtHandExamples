package main

import (
	"encoding/json"
	"errors"
)

var loginFailed = errors.New("login failed")

type argumentError struct {
	argName string
	value   string
	err     string
}

func (ae *argumentError) Error() string {
	return "Argument error:" + ae.err + ", name:" + ae.argName + ", value:" + ae.value
}

type loginData struct {
	Email    string
	Password string
}

func main() {
	inputs := []string{``,
		`{"Email":1,"Password":"b"}`,
		`{"Email":"a","Password":"c"}`,
		`{"Email":"a","Password":"b"}`,
	}

	var err error
	for _, input := range inputs {
		println("login data:", input)

		err = login(input)
		if err != nil {
			if err == loginFailed {
				println("login failed.")
			}

			switch err := err.(type) {
			case *argumentError:
				println("argument error. argument name:", err.argName)
			case *json.UnmarshalTypeError:
				println("unmarshal type error, type:", err.Value, ", expected type:", err.Type.String())
			default:
				println("other kind of error:", err.Error())
			}
		} else {
			println("login succeeded!")
		}

		println("")
	}
}

func login(jsonStr string) (err error) {
	if jsonStr == "" {
		err = &argumentError{argName: "jsonStr", value: jsonStr, err: "should not be empty"}
		return
	}

	ld := new(loginData)
	err = json.Unmarshal([]byte(jsonStr), ld)

	if err != nil {
		return
	}

	if ld.Email == "" || ld.Password == "" {
		err = errors.New("email or Password missing")
		return
	}

	//mock up the matching process
	if ld.Email == "a" && ld.Password != "b" {
		err = loginFailed
		return
	}

	return
}

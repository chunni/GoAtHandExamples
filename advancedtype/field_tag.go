package main

import (
	"encoding/json"
	"fmt"
)

type Job struct {
	Id  string `json:"id"`
	Cmd string `json:"cmd,omitempty"`
}

func marshalJob(j Job) {
	bts, err := json.Marshal(j)
	if err == nil {
		fmt.Println(string(bts))
	} else {
		fmt.Println("marshal error:", err)
	}
}

func main() {
	j1 := Job{"job 1", "command 1"}
	marshalJob(j1)

	j2 := Job{Id: "job 2"}
	marshalJob(j2)

	j3 := Job{Cmd: "command 3"}
	marshalJob(j3)
}

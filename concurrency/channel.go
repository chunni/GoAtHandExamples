package main

func main() {
	msgChan := make(chan string)

	go echo(msgChan)

	msgToSend := []string{"Hi", "Hello", "Salute"}

	for _, msg := range msgToSend {
		msgChan <- msg
		println(<-msgChan)
	}
}

func echo(queue chan string) {
	for msg := range queue {
		println(msg)
		queue <- msg + "," + msg
	}
}

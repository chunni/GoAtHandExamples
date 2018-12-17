package main

type player struct {
	name   string
	points int
}

//value method
func (t player) draw() {
	t.points += 1
}

//pointer method
func (t *player) win() {
	t.points += 3
}

func main() {
	p := player{name: "Jim"}

	p.draw()
	println(p.name+"'s points after a draw:", p.points)

	p.win()
	println(p.name+"'s points after a win:", p.points)

	pointer := &p
	pointer.draw()
	pointer.win()
}

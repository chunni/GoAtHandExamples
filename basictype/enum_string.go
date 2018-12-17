package main

type Season string

const (
	Spring Season = "Spring"
	Summer        = "Summer"
	Autumn        = "Autumn"
	Winter
)

func main() {
	printSeason(Spring)
	printSeason(Summer)
	printSeason(Autumn)
	printSeason(Winter)
}

func printSeason(season Season) {
	println("Season:", season)
}

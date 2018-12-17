package main

type day int

const (
	sunday day = iota
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

var dayString = [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

func (d day) String() string {
	return dayString[d]
}

func DayFromString(s string) day {
	for k, v := range dayString {
		if v == s {
			return day(k)
		}
	}
	return sunday
}

func main() {
	d := wednesday
	println("Value:", d, ", Text:", d.String())

	d1 := DayFromString("Friday")
	println("Value:", d1, ", Text:", d1.String())
}

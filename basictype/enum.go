package main

type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	println("Monday:", Monday)
	println("Saturday:", Saturday)

	i := 2
	println(i == int(Tuesday))

	println(Thursday.String())
}

func (day Day) String() string {
	switch day {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday:
		return "Saturday"
	}
	return ""
}

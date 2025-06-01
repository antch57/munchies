package models

type Snack struct {
	Snack string
	Count int
	Time  string
}

type Command struct {
	Name string
	Help string
	Run  func(args []string) error
}

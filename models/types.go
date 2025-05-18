package models

type Snack struct {
	Snack string
	Count int
}

type Command struct {
	Name string
	Help string
	Run  func(args []string) error
}

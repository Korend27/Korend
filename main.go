package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type worker struct {
	Info       person
	Expirience int
}

type doctor struct {
	Human worker
	Spec string
	Famdoc bool
	Backalavr bool
	Magistr bool
	Job string
}
type doctor struct {
	Human worker
	Spec string
	Famdoc bool
	Backalavr bool
	Magistr bool
	Job string
}
func main() {
	delly := doctor{
		worker{
		person{
			"Delly",
			24},
		3},
		"Физиотерапевт",
		false,
		true,
		true,
		"Приватный врач"}

	fmt.Println(delly)
}

package main

import (
	_ "net/http/pprof"
)

// комментарий к функции
func main() {
	println("Hello world !!!")
	println("5-1+0")
	println("5-1+22")
	println("5-1+22+0")
}

func second() {
	println("second")
}

package main

import (
	_ "net/http/pprof"
)

// комментарий к функции
func main() {
	println("Hello world !!!")
}

func second() {
	println("second")
}

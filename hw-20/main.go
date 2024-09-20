package main

import (
	_ "net/http/pprof"
)

// комментарий к функции
func main() {
	println("Hello world !!!")
	println("5-1+1")
}

func second() {
	println("second")
}

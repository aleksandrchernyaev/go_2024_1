package main

import (
	"fmt"
	"time"
	"github.com/beevik/ntp
)

func main() {

	t := time.Now()
	fmt.Println(t.Format(time.TimeOnly))
	fmt.Println("Hello world")
}

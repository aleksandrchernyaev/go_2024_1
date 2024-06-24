package main

import "time"

func main() {

	slice := []func() error{}

	max_number_routites := 10
	max_errors := 2
	num_errors := 0

	ch_number_routites := make(chan bool, max_number_routites)
	ch_errors := make(chan bool)

	for key := range slice {

		f := slice[key]
		_, ok := <-ch_errors
		if ok {
			num_errors++
		}
		if num_errors >= max_errors {
			break
		}

		go do_function(ch_number_routites, ch_errors, f)
		ch_number_routites <- true
	}

	println("start")

	timer := time.NewTimer(10 * time.Second)
	<-timer.C
	close(ch_number_routites)
	close(ch_errors)

	println("finish")

}

func do_function(ch_number_routites <-chan bool, ch_errors chan<- bool, f func() error) {

	er := f()
	if er != nil {
		ch_errors <- true
	}
	<-ch_number_routites

}

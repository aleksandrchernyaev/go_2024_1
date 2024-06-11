package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var txt string
	var low_word string
	dictionary := map[string]int{}

	type key_value struct {
		key   string
		value int
	}

	txt = `This string lines is on
	multiple lines
	within three single
	quotes on either lines side.`

	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%,.!_ ;:\n-?	", r)
	}

	words := strings.FieldsFunc(txt, splitFunc)
	for _, word := range words {
		low_word = strings.ToLower(word)
		//fmt.Printf("Word %d is: %s \n", idx, low_word)

		if _, ok := dictionary[low_word]; !ok {
			dictionary[low_word] = 0
		}
		dictionary[low_word] = dictionary[low_word] + 1
		//sorted_struct = append(sorted_struct , key_value {key, value})
	}
	//var slice_struct []key_value

	slice := make([]key_value, len(dictionary)+1)

	CompareFunc := func(i, j int) bool {
		return slice[i].value > slice[j].value
	}

	for key, value := range dictionary {
		slice = append(slice, key_value{key, value})
	}

	sort.Slice(slice, CompareFunc)

	i := 0
	for _, key_value := range slice {
		fmt.Printf("%s, %d \n", key_value.key, key_value.value)
		i++
		if i == 5 {
			break
		}
	}
}

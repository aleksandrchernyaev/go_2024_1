package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	type pair struct {
		s1 string
		s2 string
	}

	test := []pair{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{`qwe\4\5`, `qwe45`},
		{`qwe\45`, `qwe44444`},
		{`qwe\\5`, `qwe\\\\\`},
	}

	for _, t := range test {

		str, _ := decode(t.s1)
		if t.s2 == str {
			fmt.Printf("%s - %s\n", t.s1, "OK")
		} else {
			fmt.Printf("%s - %s\n", t.s1, "FAIL")
		}
	}

}

func decode(s string) (string, error) {
	var num int
	var err1 error
	var err2 error

	r := []rune(s)
	s1 := ""
	str := ""
	show := false
	escape := false

	for i := 0; i < len(r); i++ {

		num, err1 = strconv.Atoi(string(r[i]))

		if (s1 != "\\" || escape) && err1 == nil {
			//удалось преобразовать в число
			if show {
				for j := 0; j < num; j++ {
					str = str + s1
				}
			} else {
				err2 = errors.New("Не удалось преобразовать строку")
				return "", err2
			}
			show = false //не нужно выводить цифру
		} else {

			if s1 != "\\" && show {
				str = str + s1
			}
			show = true //нужно выводить символ
		}

		if string(r[i]) == "\\" && s1 == "\\" {
			escape = true
		}

		s1 = string(r[i])
	}

	if show {
		str = str + s1
	}

	return str, nil
}

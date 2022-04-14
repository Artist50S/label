package main

import (
	"fmt"
	"regexp"
)

type badLabel struct {
	rus          bool
	numberSimvol bool
	link         bool
}

func main() {
	label := "qввq http//12.ru"
	bl := NewbadLabel(label)
	fmt.Println(bl)
}

//true - недопустимое значение, false - good label
func NewbadLabel(lb string) badLabel {
	var bl badLabel
	bl.rus = lat(lb)
	bl.numberSimvol = numberSimvol(lb)
	bl.link = link(lb)
	return bl
}

func lat(str string) bool {
	rus, err := regexp.MatchString("[А-Яа-я]", str)
	if err != nil {
		fmt.Println(err)
	}
	return rus
}

func numberSimvol(str string) bool {
	s := len(str)
	if s > 20 {
		nS := true
		return nS
	} else {
		nS := false
		return nS
	}
}

func link(str string) bool {
	re, err := regexp.Compile("http://|https://|.ru|.com|.dev")
	http := re.MatchString(str)
	if err != nil {
		fmt.Println(err)
	}
	return http
}

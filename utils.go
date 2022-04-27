package main

import (
	"fmt"
	"strings"
)

func IsNullorWhitespace(str string) bool {
	var trimmed = strings.Trim(str, " ")
	if trimmed == "" {
		return true
	} else {
		return false
	}
}

func IsaNumber(str interface{}) bool {
	var k bool
	defer func() {
		if err := recover(); err != nil {
			k = false
			fmt.Println(err, k)
		}
	}()
	_ = str.(float64)
	k = true
	return k
}

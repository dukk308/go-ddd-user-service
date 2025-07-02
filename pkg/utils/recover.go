package utils

import "fmt"

func Recover() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := []rune("rêsumê")
	indexed := myString[0]
	fmt.Printf("%v, %T", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	var myRune = 'a'
	fmt.Printf("\nyRune = %v", myRune)

	var strSlice = []string{"s", "u", "b"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	catStr := strBuilder.String()
	fmt.Printf("\n%v", catStr)
}

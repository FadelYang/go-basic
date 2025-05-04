// Go data type
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum1 int = 32767
	intNum1 = intNum1 + 1
	fmt.Println(intNum1)

	var floatNum1 float32 = 12345678.9
	var floatNum2 float64 = 12345678.9
	fmt.Println(floatNum1)
	fmt.Println(floatNum2)

	var intNum3 int = 2
	var floatNum3 float64 = 21.1
	var result float64 = float64(intNum3) + floatNum3
	fmt.Println(result)

	var myString = `Hello
	Go`
	var myString2 = "Hello\nGo"
	fmt.Println(myString)
	fmt.Println(myString2)
	fmt.Println(utf8.RuneCountInString((myString)))
	fmt.Println(utf8.RuneCountInString((myString2)))

	var myVar = "Fadel"
	fmt.Println(myVar)

	myVar2 := "Fadel shorthand"
	fmt.Println(myVar2)

	numberShortHand1, numberShorthand2 := 12, 2.1
	fmt.Println(numberShortHand1, numberShorthand2)
}

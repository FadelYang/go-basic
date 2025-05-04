// Array in Go
package main

import "fmt"

func main() {
	var intArr [3]int32
	intArr[1] = 10
	fmt.Println(intArr)
	fmt.Println(intArr[1:3])

	var intArr1 [3]int32 = [3]int32{1, 10, 5}
	fmt.Println(intArr1)

	intArr2 := [...]int32{10, 12, 10}
	fmt.Println(intArr2)
}

// Array in Go
package main

import "fmt"

func main() {
	// Array dengan size yang immutable atau fixed
	var intArr [3]int32
	intArr[1] = 10
	fmt.Println(intArr)
	fmt.Println(intArr[1:3])

	var intArr1 [3]int32 = [3]int32{1, 10, 5}
	fmt.Println(intArr1)

	intArr2 := [...]int32{10, 12, 10}
	fmt.Println(intArr2)

	// Item dalam array disimpan dimemori yang
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// Slices, Array dengan size yang lebih fleksibel
	var intSlice []int32 = []int32{1, 2, 3}
	intSlice = append(intSlice, 32)
	intSlice2 := []int32{12, 2, 4}
	fmt.Println(intSlice)
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// Cara lain untuk membuat slice
	var intSlice3 []int32 = make([]int32, 2)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	var myMap2 = map[string]uint8{"Adam": 32, "Eve": 28}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])

	var data, ok = myMap2["Dinda"]
	if ok {
		fmt.Println(data)
	} else {
		fmt.Println("Data not found")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n\n", name, age)
	}

	// for lopp dan range
	persons := []string{"Miaw", "Augh", "Siapa", "Dia"}
	for i := 0; i < len(persons); i++ {
		fmt.Println(persons[i])
	}

	fmt.Println("\n")

	for _, value := range persons {
		fmt.Println("Value:", value)
	}

	fmt.Println("\n")

	for index := range persons {
		fmt.Println("Value2:", persons[index])
	}
}

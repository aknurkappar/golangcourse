package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

func main() {
	fmt.Println("Hello Go!")

	// var foo int
	// var foo int = 42
	// var foo, bar int = 42, 34
	// var foo = 42
	// foo := 32
	const constant = "This is a constant"

	// functions
	fmt.Println(functionName(34))
	fmt.Println(add(30, 10))
	fmt.Println(returnMulti())
	var x, str = returnMulti2()
	fmt.Println(x, str)

	addNumbers := func(a, b int) int {
		return a + b
	}
	fmt.Println(addNumbers(5, 6))

	// arrays
	var a [10]int
	a[3] = 10
	i := a[3]
	fmt.Println("Array", a)
	fmt.Println("3th element", i)

	// var a2 = [2]int{1, 2}
	// a2 := [2]int{1, 2}
	a2 := [...]int{1, 2}
	fmt.Println("Initialized array", a2)

	// slices
	var slice []int
	fmt.Println("Slice", slice)

	var slice2 = []int{1, 2, 3, 4}
	fmt.Println("Slice", slice2)

	chars := []string{0: "a", 2: "c", 1: "b"}
	fmt.Println("Slice", chars)

	slice3 := append(slice2, 5)
	fmt.Println("Slice", slice3)

	slice4 := slice3[1:3] // indexes 1, 2 (3 not included)
	fmt.Println("Slice", slice4)

	var sliceMake = make([]byte, 5, 5)
	fmt.Println("Slice using make", sliceMake)

	for i, e := range slice2 {
		fmt.Println("index:", i, ", element:", e)
	}
	for _, e := range slice2 {
		fmt.Println("element:", e)
	}

	// for range time.Tick(time.Second) {
	// 	fmt.Println("every 1 sec")
	// }

	// maps
	m := make(map[string]int)
	m["key"] = 42
	fmt.Println(m["key"])

	delete(m, "key")

	elem, ok := m["key"]
	if ok {
		fmt.Println("Key found:", elem, ok)
	} else {
		fmt.Println("Key not found", ok)
	}

	var locations = map[string]Vertex{
		"Bell Labs": {40.38748, -74.435454},
		"Google":    {37.54664, -122.464647},
	}

	for key, value := range locations {
		fmt.Println(key, value.Lat, value.Long)
	}
}

func functionName(param int) int {
	return param
}

func add(param1, param2 int) int {
	return param1 + param2
}

func returnMulti() (int, string) {
	return 10, "foobar"
}

func returnMulti2() (n int, s string) {
	n = 20
	s = "Aknur"
	return n, s
}

func scope() func() int {
	outer_var := 2
	foo := func() int {
		return outer_var
	}
	return foo
}

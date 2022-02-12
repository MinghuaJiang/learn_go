package main

import "fmt"

func main() {
	fmt.Println(testDefer())

}

func testDefer() int {
	i := 2
	defer func() {
		fmt.Println("test defer")
		i = 3
	}()
	return i
}

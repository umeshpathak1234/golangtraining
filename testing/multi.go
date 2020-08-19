package testing

import "fmt"

func multi(a int) int {

	return (a - 20)
}

func Testingagain(s string) chan int {
	fmt.Println("go checking")
	c := make(chan int)

	for _, v := range s {
		c <- int(v)
	}
	return c
}

package testing

import "fmt"

func Newfunc(a, b int) int {
	fmt.Println(vari("James", 2, 3, 4))
	fmt.Println(vari("Bond", 5, 4, 3, 2))
	c := (a + 5) * (b + 5)
	return (multi(c))
}

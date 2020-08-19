package testing

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type Personal func(Teststruct []person)

var P1 Personal

// Teststruct will return slice of person
func Teststruct() []person {

	p1 := person{"James", "Bond", 27}
	p2 := person{"Rocky", "Mondy", 45}
	a := []person{p1, p2}
	return a
}

func Multifunc(a Personal) {
	fmt.Println("There is nothing to print :(")
	fmt.Println(a)
}

// Teststructone will  take three input and return slice of person
func Teststructone(a string, b string, c int) []person {
	p1 := person{a, b, c}
	p2 := person{a + "y", b + "rd", c + 5}
	return []person{p1, p2}
}

//variadic fucntion which takes one string and any number of integers
// returns the length of the string and sum of integers it receives
func vari(a string, b ...int) (string, int, int) {

	var total int
	for _, v := range b {
		total += v
	}

	return "the number of alphabet in " + a + " is", len(a), total

}

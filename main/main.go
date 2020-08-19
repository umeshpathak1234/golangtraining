package main

import (
	"fmt"
)

func main() {

	// wg.Add(1)

	// fmt.Println(a)
	a := make(chan string)

	go Testingagaina(a)
	v1 := <-a
	fmt.Println(v1)

	// wg.Wait()

	// fmt.Println("The output is", testing.Newfunc(2, 3))
	// fmt.Println("The new structures created are", testing.Teststruct())
	// fmt.Println("NEw function", testing.Teststructone("Rock", "wa", 27))
	// testing.Multifunc(testing.P1)

	// interfaces.Info(interfaces.Square{10})
	// interfaces.Info(interfaces.Circle{10})
}

func Testingagaina(c chan string) chan string {

	//c <- someval + "hi there"
	//t := map[string]string{"a": "apple", "b": "banana"}
	t := []string{"1", "b", "ch", "d", "e"}
	fmt.Println(t)
	for i, v := range t {
		fmt.Println(i, v)
		c <- (v)
		fmt.Println(v, i)
	}
	return c
}

type umesh struct {
	Name string
}
type tert interface {
	abc()
}

func abc(s int) float64 {
	return float64(s) * 5.0
}
func abhc(u tert) {
	fmt.Println("u")
}
func (t umesh) abc() string {

	return t.Name
}

type pathak struct {
	lname string
}

func (u pathak) abc() string {
	Name := "Pathak"
	return Name
}

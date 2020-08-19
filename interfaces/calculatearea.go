package interfaces

//interfaces are type
import (
	"fmt"
	"math"
)

type Square struct {
	Side float64
}
type Circle struct {
	Radius float64
}

// square implements the shape interface
func (s Square) area() float64 {
	fmt.Println("The area of Square is")
	return s.Side * s.Side
}
func (s Circle) area() float64 {
	fmt.Println("The area of Circle is")
	return math.Pi * s.Radius * s.Radius
}

/* interfaces are types and can be defined how that
can be implemented
in this condition anything that has the method are() float64 is
implementing the Shape method
*/
type Shape interface {
	area() float64
}

func Info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

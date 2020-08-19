package main

import (
	"fmt"
)

type Company interface {
	moneySource() string
	totalMoney() float64
}

type Billing struct {
	companyName    string
	directExpenses float64
}
type Salary struct {
	noOfHours   int
	ratePerHour float64
}

func (bm Salary) totalMoney() float64 {

	return bm.ratePerHour * float64(bm.noOfHours)
}
func (bm Billing) moneySource() string {
	return bm.companyName
}

func (sa Billing) totalMoney() float64 {
	return sa.directExpenses
}
func (sa Salary) moneySource() string {
	return "Name"
}
func calculatetotalexpenses(ac []Company) {
	var netincome float64
	for _, totalexpenses := range ac {

		fmt.Println("The source of the money is", totalexpenses.moneySource())
		fmt.Println("The total amount of money is ", totalexpenses.totalMoney())
		netincome += totalexpenses.totalMoney()
	}
	fmt.Println("The total income of the is", netincome)

}

// //func Newfunc(a, b int) int {
// 	fmt.Println(vari("James", 2, 3, 4))
// 	fmt.Println(vari("Bond", 5, 4, 3, 2))
// 	c := (a + 5) * (b + 5)
// 	return (multi(c))
// }
func main() {
	pa := Billing{"Google", 5000.987}
	pa1 := Salary{160, 55.35}
	pa3 := []Company{pa, pa1}
	//incomeStreams := []income{project1, project2, project3, project4}
	calculatetotalexpenses(pa3)
	//calculatetotalexpenses(pa1)

}

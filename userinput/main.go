package main

import (
	"fmt"
)

func main() {
	fmt.Println("Testing")
	// to run this Testing() function both the main.go and test.go must be run
	// go run *.go which will build and run all the .go files in the folder main package
	// or use go build and it will build a excutable file of the same name of the folder
	// in this case executable  file will be userinput. double click that file it will run
	//Testing()

	/*
		var names = []string{"Umesh", "zebra", "van"}

		fmt.Println("The slice is as it is", names)
		//isSiceSortedOne := sort.StringsAreSorted(names)
		if sort.StringsAreSorted(names) {
			fmt.Println((names[0]))
		} else {
			fmt.Println(names[2])
		}
		sort.Strings(names)
		fmt.Println("The slice is sorted", names)

		isSiceSorted := sort.StringsAreSorted(names)
		if isSiceSorted == true {
			fmt.Println((names[0]))
		} else {
			fmt.Println(names[2])
		}
		sort.Sort(sort.Reverse(sort.StringSlice(names)))
		fmt.Println("The slice is now reverse sorted", names)
	*/
	power := 1000
	fmt.Printf("Default power is %d\n", power)
	name, power := "Umesh", 32
	fmt.Printf("%s's power is %d\n", name, power)

}

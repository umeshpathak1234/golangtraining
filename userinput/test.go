package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Testing() {
	fmt.Println("testing successfull")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Plese type your full name")
	name, _ := reader.ReadString('\n')
	fmt.Println("The name that  you entered is\t", name)

	ratingreader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your rating")
	review, _ := ratingreader.ReadString('\n')
	numrating, _ := strconv.ParseFloat(strings.TrimSpace(review), 64)
	// converting user input string review to float64

	fmt.Println("Your entered reveiew is", numrating+2)
}

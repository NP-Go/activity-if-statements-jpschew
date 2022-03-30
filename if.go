package main

import (
	"fmt"
	"strconv"
)

func main() {
	// number1 := 17
	number2 := 24
	resultMessage := "No outcome."
	//Insert your code here
	//Hint: You may wish to make use of strconv.Itoa to convert int to string

	// using one preceding condition
	if number1 := 24; number1 > number2 {
		resultMessage = strconv.Itoa(number1) + " is bigger than " + strconv.Itoa(number2) + " by " + strconv.Itoa(number1-number2)
		fmt.Println(resultMessage)
	} else if number2 > number1 {
		resultMessage = strconv.Itoa(number2) + " is bigger than " + strconv.Itoa(number1) + " by " + strconv.Itoa(number2-number1)
		fmt.Println(resultMessage)
	} else {
		fmt.Println(resultMessage)
	}
}

package main

import (
	"bufio";
	"fmt";
	"os";
	"strconv";
	"strings"
)

func main() {

	welcome := "Welcome to the conversion program!"
	fmt.Println(welcome)

	fmt.Println("Enter the serial number")
    reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for entering the serial number: ", input)

	rollNum, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println("The roll number is: ", rollNum+1)
	}
}
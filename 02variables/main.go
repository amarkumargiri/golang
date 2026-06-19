package main

import "fmt"

func main(){
	var username string = "golang"
	fmt.Println("Hello", username)          //println prints with a new line
	fmt.Printf("variable is of type : %T\n", username) //printf  prints without a new line and allows formatting

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)          
	fmt.Printf("variable is of type : %T\n", isLoggedIn) 

	var smallVal uint8 = 255  // unit8 can only store values from 0 to 255 
	fmt.Println(smallVal)          
	fmt.Printf("variable is of type : %T\n", smallVal) 

	var smallFloat float32 = 255.4555555555 // float32 can only store values up to 5 places
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type : %T\n", smallFloat)
}
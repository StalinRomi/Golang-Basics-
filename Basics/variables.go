package main

import "fmt"

func main() {
	var username string = "Stalin"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// implicit type declaration

	var website = "www.google.com"
	fmt.Println(website) 


}
package main

import (
	"fmt"
	"os"
)

func User_input() {
	message := "Welcome to User Input"
	fmt.Println(message)

	reader := buffio.NewReader(os.Stdin)
	fmt.Println("Enter the number")
	
}
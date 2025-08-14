package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)
	userNames[0] = "John"
	userNames[1] = "Jane"
	userNames = append(userNames, "Chevy")
	userNames = append(userNames, "Ram")
	fmt.Println(userNames)
}

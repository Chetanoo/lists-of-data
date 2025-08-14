package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)
	userNames[0] = "John"
	userNames[1] = "Jane"
	userNames = append(userNames, "Chevy")
	userNames = append(userNames, "Ram")
	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 3)
	courseRatings["Go"] = 4.5
	courseRatings["Python"] = 4.7
	courseRatings["Java"] = 4.9
	fmt.Println(courseRatings)
}

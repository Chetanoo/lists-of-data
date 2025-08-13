package main

import "fmt"

func main() {
	// 1
	hobbies := [3]string{"Reading", "Gaming", "Walking"}
	fmt.Println(hobbies)

	// 2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// 3
	firstSlice := hobbies[:2]
	secondSlice := firstSlice[:2]
	fmt.Println(firstSlice, secondSlice)

	// 4
	thirdSlice := firstSlice[1:3]
	fmt.Println(thirdSlice)

	// 5
	goals := []string{"learn", "improve attention span"}

	// 6
	goals[1] = "different"
	goals = append(goals, "new goal")
	fmt.Println("goals", goals)

	// 7
	type Product struct {
		title string
		id    string
		price float64
	}

	products := []Product{
		{
			title: "Phone",
			id:    "1",
			price: 999.99,
		},
		{
			title: "Tablet",
			id:    "2",
			price: 699.99,
		},
	}
	fmt.Println(products)
	products = append(products, Product{
		title: "Headphones",
		id:    "3",
		price: 399.99,
	})
	fmt.Println(products)
}

//func main() {
//	prices := []float64{10.99, 8.99}
//	fmt.Println(prices[0])
//
//	updatedPrices := append(prices, 13.99)
//	fmt.Println(updatedPrices)
//}

//func main() {
//	var productNames [4]string = [4]string{"A Book"}
//	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
//	fmt.Println(prices)
//
//	productNames[2] = "A Carpet"
//
//	fmt.Println(productNames)
//
//	fmt.Println(prices[2])
//
//	slice := prices[1:3]
//	featuredPrices := prices[1:]
//	highlightedPrices := featuredPrices[:1]
//	fmt.Println(featuredPrices)
//	fmt.Println(len(featuredPrices))
//	fmt.Println(cap(featuredPrices))
//	fmt.Println(highlightedPrices)
//	fmt.Println(slice)
//
//	highlightedPrices = highlightedPrices[:3]
//	fmt.Println(highlightedPrices) // this is a weird behavior but it works like this in go.
//}

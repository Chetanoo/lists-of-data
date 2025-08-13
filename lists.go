package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0])

	updatedPrices := append(prices, 13.99)
	fmt.Println(updatedPrices)
}

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

package main

import "fmt"

type Product struct {
	id    int
	name  string
	price float64
}

type TemperatureData struct {
	temperature float64
	humidity    float64
}

func arrays_slices() {
	var productNames [4]string
	prices := [6]float64{1.2, 3.4, 5.6, 7.8, 9.0, 10.2}
	fmt.Println(prices)

	productNames[0] = "Apple"
	productNames[1] = "Banana"
	productNames[2] = "Cherry"
	fmt.Println(productNames[2])

	featuredPrices := prices[1:4]
	fmt.Println(featuredPrices)

	featuredPrices = prices[:]
	fmt.Println(featuredPrices)

	fmt.Println(len(featuredPrices), cap(featuredPrices))

	discountPrices := append(featuredPrices, prices[4:]...)
	fmt.Println(discountPrices)
}

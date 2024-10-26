package main

import (
	"fmt"
	"price_calculator/filemanager"
	"price_calculator/prices"
)

func main() {
	taxRates := []float64{0.00, 0.08, 0.10, 0.20}

	result := make(map[float64][]float64)

	for _, taxRates := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%v.json", taxRates*100))
		//cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(taxRates, fm)
		err := priceJob.Process()
		if err != nil {
			fmt.Println(err)
		}

	}

	fmt.Println(result)
}

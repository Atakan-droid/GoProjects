package prices

import (
	"errors"
	"fmt"
	"price_calculator/conversion"
	"price_calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"` // Ignore this field in JSON
}

// Constructor
func NewTaxIncludedPriceJob(taxRate float64, io iomanager.IOManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   io,
	}
}

func (job *TaxIncludedPriceJob) Process() error {
	//Load data
	err := job.LoadData()
	if err != nil {
		return err
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	return job.IOManager.WriteResult(job)
}

func (job *TaxIncludedPriceJob) LoadData() error {
	//Read lines from file
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println("Error reading file", err)
		return errors.New("error reading file")
	}
	//Convert strings to floats
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return errors.New("error converting strings to floats")
	}

	job.InputPrices = prices
	return nil
}

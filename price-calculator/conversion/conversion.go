package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floatList []float64

	for _, line := range strings {
		floatVal, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, errors.New("error converting string to float")
		}
		floatList = append(floatList, floatVal)
	}

	return floatList, nil
}

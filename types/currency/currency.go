package currency

import (
	"fmt"
)

// Maximum 99999999
type Amount float64

// ex. return 999,999.99
func (c Amount) ToString() (result string) {
	amount := fmt.Sprintf("%.2f", c/100)
	splitter := len(amount) - 3
	result = amount[splitter+1:]

	for i := splitter; i >= 0; i-- {
		if (splitter-i)%4 == 0 && splitter != i {
			result = fmt.Sprintf("%s%s", ",", result)
		}
		result = fmt.Sprintf("%c%s", amount[i], result)
	}

	offset := 10 - len(result)
	for i := 0; i < offset; i++ {
		result = fmt.Sprintf("%s%s", " ", result)
	}

	return result
}

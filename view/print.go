package view

import (
	"fmt"
	"go-tamboon/types/currency"
)

type DonationResult struct {
	TotalReceived       float64
	SuccessfullyDonated float64
	FaultyDonate        float64
	AvgPerPerson        float64
	TopDonate           []string
}

func (dr DonationResult) Print() {
	// fmt.Println("let's print result")
	tmp := `
              total received: THB  %s
        successfully donated: THB  %s
             faulty donation: THB  %s
  
          average per person: THB  %s
                  top donors: `

	topDonateLimit := 0
	if len(dr.TopDonate) > 3 {
		topDonateLimit = 3
	} else {
		topDonateLimit = len(dr.TopDonate)
	}

	for i := 0; i < topDonateLimit; i++ {
		if i == 0 {
			tmp = fmt.Sprintf("%s%s\n", tmp, dr.TopDonate[i])
			continue
		}

		tmp = fmt.Sprintf("%s                              %s\n", tmp, dr.TopDonate[i])
	}

	fmt.Printf(tmp,
		currency.Amount(dr.TotalReceived).ToString(),
		currency.Amount(dr.SuccessfullyDonated).ToString(),
		currency.Amount(dr.FaultyDonate).ToString(),
		currency.Amount(dr.AvgPerPerson).ToString())

}

package main

import (
	"encoding/csv"
	"fmt"
	"go-tamboon/controllers"
	"go-tamboon/models"
	"go-tamboon/queues"
	"go-tamboon/view"
	"os"
	"strconv"
)

func main() {
	// dat, _ := os.ReadFile("data/fng.1000.csv.rot128")
	// // r := strings.NewReader(string(dat))
	// rot128reader, err := cipher.NewRot128Reader(bytes.NewBuffer(readDat()))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// buff := make([]byte, len(dat))
	// rot128reader.Read(buff)

	csvFile, err := os.Open("data/fng.1000.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	usingCsv := csvLines[1:]

	queues.CreateChannel()
	go queues.Do()
	fmt.Println("performing donations...")

	for _, line := range usingCsv {
		ci := models.ChargeInfo{
			Name:     line[0],
			Amount:   line[1],
			CCNumber: line[2],
			CVV:      line[3],
			ExpMonth: line[4],
			ExpYear:  line[5],
		}
		if err := controllers.LetsDonation(ci); err != nil {
			amount, err := strconv.ParseFloat(ci.Amount, 64)
			if err == nil {
				go queues.AddFault(amount)
			}
		} else {
			go queues.AddCI(ci)
		}
	}
	fmt.Println("Done.")
	go queues.Exit()

	donationResult := view.DonationResult{
		TotalReceived:       queues.TotalReceived,
		SuccessfullyDonated: queues.TotalReceived - queues.FaultyDonate,
		FaultyDonate:        queues.FaultyDonate,
		AvgPerPerson:        queues.TotalReceived / float64(len(queues.Top3Donate)),
		TopDonate:           queues.Top3Donate.ToString(),
	}

	donationResult.Print()
}

// func readDat() (result []byte) {
// 	client := &http.Client{}
// 	resp, err := client.Get("https://raw.githubusercontent.com/omise/challenges/challenge-go/data/fng.1000.csv.rot128")
// 	if err != nil {
// 		return result
// 	}
// 	defer resp.Body.Close()

// 	result, err = ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(string(result))
// 	return result
// }

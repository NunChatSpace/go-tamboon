package queues

import (
	"fmt"
	"go-tamboon/models"
	"go-tamboon/types/list"
	"strconv"
)

var CIs chan models.ChargeInfo
var FaultDonate chan float64
var Quit chan int

var Top3Donate list.Top3Rank
var AvgPerPerson float64
var TotalReceived float64
var FaultyDonate float64

func AddCI(ci models.ChargeInfo) {
	CIs <- ci
}

func AddFault(amount float64) {
	FaultDonate <- amount
}

func CreateChannel() {
	CIs = make(chan models.ChargeInfo, 10)
	FaultDonate = make(chan float64)
	Quit = make(chan int)
}

func Exit() {
	Quit <- 0
}

func Do() {
	for {
		select {
		case ci := <-CIs:
			updateSuccessInfo(ci)
		case fault := <-FaultDonate:
			updateFaultDonate(fault)
		case <-Quit:
			fmt.Println("quit")
			// close(CIs)
			return
		}
	}

}

func updateSuccessInfo(ci models.ChargeInfo) {
	amount, err := strconv.ParseFloat(ci.Amount, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	Top3Donate = Top3Donate.Add(list.Donater{
		Name:   ci.Name,
		Amount: amount / 100,
	})

	TotalReceived = TotalReceived + amount
}

func updateFaultDonate(donate float64) {
	TotalReceived = TotalReceived + (donate / 100)
	FaultyDonate = FaultyDonate + (donate / 100)
}

package models

type CustomerPayment struct {
	Name   string
	Amount string
}

type ChargeInfo struct {
	Name     string
	Amount   string
	CCNumber string
	CVV      string
	ExpMonth string
	ExpYear  string
}

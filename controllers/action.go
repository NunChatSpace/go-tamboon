package controllers

import (
	"fmt"
	"go-tamboon/models"
	"go-tamboon/sdk_omise/card"
	"go-tamboon/sdk_omise/customer"
	"go-tamboon/sdk_omise/payment"
)

func LetsDonation(ci models.ChargeInfo) error {
	addCard := card.Add{
		Name:            ci.Name,
		Number:          ci.CCNumber,
		SecurityCode:    ci.CVV,
		ExpirationMonth: ci.ExpMonth,
		ExpirationYear:  ci.ExpYear,
	}

	addCardResult, err := addCard.Execute()
	if err != nil {
		return err
	}

	addCustomer := customer.Add{
		CardToken: fmt.Sprintf("%v", addCardResult["id"]),
	}

	addCusResult, err := addCustomer.Execute()
	if err != nil {

		return err
	}

	charge := payment.Charge{
		Amount:     ci.Amount,
		Currency:   "thb",
		CustomerID: fmt.Sprintf("%v", addCusResult["id"]),
	}

	_, err = charge.Execute()
	if err != nil {
		return err
	}

	return nil
}

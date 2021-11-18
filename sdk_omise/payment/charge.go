package payment

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Charge struct {
	Amount     string `json:"amount"`
	Currency   string `json:"currency"`
	CustomerID string `json:"customer_id"`
}

func (c Charge) Execute() (result map[string]interface{}, err error) {
	params := url.Values{}
	params.Add("amount", c.Amount)
	params.Add("currency", c.Currency)
	params.Add("customer", c.CustomerID)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://api.omise.co/charges", body)
	if err != nil {
		return result, err
	}
	req.SetBasicAuth("skey_test_5pq5kcxu1j451eiw0uy", "")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("%s", resp.Status)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal([]byte(bodyBytes), &result)

	return result, nil
}

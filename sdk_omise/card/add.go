package card

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Add struct {
	Name            string `json:"name"`
	Number          string `json:"number"`
	SecurityCode    string `json:"security_code"`
	ExpirationMonth string `json:"expiration_month"`
	ExpirationYear  string `json:"expiration_year"`
}

func (c Add) Execute() (result map[string]interface{}, err error) {
	params := url.Values{}
	params.Add("card[name]", c.Name)
	params.Add("card[number]", c.Number)
	params.Add("card[security_code]", c.SecurityCode)
	params.Add("card[expiration_month]", c.ExpirationMonth)
	params.Add("card[expiration_year]", c.ExpirationYear)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://vault.omise.co/tokens", body)
	if err != nil {
		return result, err
	}

	req.SetBasicAuth("pkey_test_5pq5kcxtt5y73qgh6pt", "")
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
		return result, err
	}
	json.Unmarshal([]byte(bodyBytes), &result)

	return result, nil
}

package customer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Add struct {
	CardToken string `json:"token"`
}

func (c Add) Execute() (result map[string]interface{}, err error) {
	params := url.Values{}
	params.Add("card", c.CardToken)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://api.omise.co/customers", body)
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

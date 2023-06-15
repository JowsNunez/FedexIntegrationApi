package handlers

import (
	"bytes"
	"encoding/json"
	"fedex/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	token := GetToken()
	payload := models.Payload{
		Account: models.Account{Value: "510087020"},
		RequestedShipment: models.RequestedShipment{
			Recipient: models.Recipient{
				Address: models.Address{
					CountryCode: "MX",
					PostalCode:  83000}},
			Shipper: models.Shipper{
				Address: models.Address{
					CountryCode: "MX",
					PostalCode:  85000}},
			PickupType:        "DROPOFF_AT_FEDEX_LOCATION",
			RateRequestType:   []string{"PREFERRED"},
			ServiceType:       "FEDEX_EXPRESS_SAVER",
			PackagingType:     "YOUR_PACKAGING",
			PreferredCurrency: "MXN",
			ServiceTypeDetail: models.ServiceTypeDetail{
				CarrierCode: "FDXE",
			},
			RequestPackageLines: []models.RequestPackageLines{
				{
					Weight: models.Weight{
						Units: "KG",
						Value: 1,
					},
				},
			},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	req, err := http.NewRequest("POST", "https://apis-sandbox.fedex.com/rate/v1/rates/quotes", bytes.NewBuffer(jsonPayload))

	if err != nil {
		log.Fatal("An error occurred from request", err)
	}
	defer req.Body.Close()

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal("An error occurred from response", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("An error occurred from get body", err)
	}
	var sm models.ResponseRate
	json.Unmarshal([]byte(body), &sm)
	json, err := json.Marshal(sm)
	fmt.Fprintln(w, string(json))
}

func GetToken() string {

	form := url.Values{}
	form.Add("grant_type", LoadValue("GRANT_TYPE"))
	form.Add("client_id", LoadValue("CLIENT_ID"))
	form.Add("client_secret", LoadValue("CLIENT_SECRET"))
	req, err := http.NewRequest("POST", "https://apis-sandbox.fedex.com/oauth/token", strings.NewReader(form.Encode()))

	if err != nil {
		log.Fatal("An error occurred request", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal("An error occurred from get response", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var TokenResp models.TokenResp
	json.Unmarshal([]byte(body), &TokenResp)
	return TokenResp.AccessToken

}

func LoadValue(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("An error occurred from load enviroment variables", err)
	}

	value := os.Getenv(key)
	fmt.Println(value)
	return value
}

package models

type Payload struct {
	Account           Account           `json:"accountNumber"`
	RequestedShipment RequestedShipment `json:"requestedShipment"`
}
type Account struct {
	Value string `json:"value"`
}
type Shipper struct {
	Address Address `json:"address"`
}
type Recipient struct {
	Address Address `json:"address"`
}

type RequestedShipment struct {
	Shipper             Shipper               `json:"shipper"`
	Recipient           Recipient             `json:"recipient"`
	PickupType          string                `json:"pickupType"`
	ServiceType         string                `json:"serviceType"`
	PackagingType       string                `json:"packagingType"`
	RateRequestType     []string              `json:"rateRequestType"`
	PreferredCurrency   string                `json:"preferredCurrency"`
	ServiceTypeDetail   ServiceTypeDetail     `json:"serviceTypeDetail"`
	RequestPackageLines []RequestPackageLines `json:"requestedPackageLineItems"`
}
type Address struct {
	PostalCode  int    `json:"postalCode"`
	CountryCode string `json:"countryCode"`
}

type ServiceTypeDetail struct {
	CarrierCode string `json:"carrierCode"`
}
type RequestPackageLines struct {
	Weight Weight `json:"weight"`
}

type Weight struct {
	Units string `json:"units"`
	Value int32  `json:"value"`
}
type TokenResp struct {
	AccessToken string `json:"access_token"`
}

type ResponseRate struct {
	OutputRate OutputRate `json:"output"`
}
type OutputRate struct {
	RateReplyDetails []RateReplyDetails `json:"rateReplyDetails"`
}
type RateReplyDetails struct {
	RatedShipmentDetails []RatedShipmentDetails `json:"ratedShipmentDetails"`
}
type RatedShipmentDetails struct {
	TotalNetCharge      float32 `json:"totalNetCharge"`
	TotalBaseCharge     float32 `json:"totalBaseCharge"`
	TotalNetFedExCharge float32 `json:"totalNetFedExCharge"`
}

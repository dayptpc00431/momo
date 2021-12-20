package models

//define a POS payload, reference in https://developers.momo.vn/#thanh-toan-pos-xu-ly-thanh-toan
type PosHash struct {
	PartnerCode  string `json:"partnerCode"`
	PartnerRefID string `json:"partnerRefId"`
	Amount       int    `json:"amount"`
	PaymentCode  string `json:"paymentCode"`
}

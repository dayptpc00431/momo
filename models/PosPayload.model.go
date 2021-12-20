package models

//define a payload, reference in https://developers.momo.vn/#cong-thanh-toan-momo-phuong-thuc-thanh-toan
type PosPayload struct {
	PartnerCode  string `json:"partnerCode"`
	PartnerRefID string `json:"partnerRefId"`
	Hash         string `json:"hash"`
	Version      int    `json:"version"`
}

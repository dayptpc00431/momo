package services

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/dayptpc00431/Momo-Payment/commons"
	"github.com/dayptpc00431/Momo-Payment/models"
	"github.com/google/uuid"
)

var GenerateURL = func(dto commons.RequestDTO) interface{} {
	requestId := uuid.New()
	var mapData = map[string]string{
		"accessKey":   dto.AccessKey,
		"amount":      strconv.FormatUint(dto.Amount, 10),
		"extraData":   commons.ExtraData,
		"ipnUrl":      commons.IpnUrl,
		"orderId":     dto.OrderID,
		"orderInfo":   dto.OrderInfo,
		"partnerCode": dto.PartnerCode,
		"redirectUrl": dto.RedirectURL,
		"requestId":   requestId.String(),
		"requestType": commons.RequestType,
		"secretKey":   dto.SecretKey,
	}

	signature := commons.RawSignature(mapData)

	var payload = models.Payload{
		PartnerCode: dto.PartnerCode,
		AccessKey:   dto.AccessKey,
		RequestID:   requestId.String(),
		Amount:      strconv.FormatUint(dto.Amount, 10),
		OrderID:     dto.OrderID,
		OrderInfo:   dto.OrderInfo,
		RedirectUrl: dto.RedirectURL,
		IpnUrl:      commons.IpnUrl,
		ExtraData:   commons.ExtraData,
		RequestType: commons.RequestType,
		Signature:   signature,
	}

	//send HTTP to momo endpoint
	resp, _, erro := commons.HttpPost(commons.Endpoint, payload).BaseRequest.End()
	if erro != nil {
		fmt.Println(erro)
	}
	fmt.Println("BODY: ", resp.Body)
	//result
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println("Response from Momo: ", result)

	return result
}

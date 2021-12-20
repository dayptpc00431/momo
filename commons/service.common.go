package commons

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var RawSignature = func(data map[string]string) string {
	var rawSignature bytes.Buffer
	rawSignature.WriteString("accessKey=")
	rawSignature.WriteString(data["accessKey"])
	rawSignature.WriteString("&amount=")
	rawSignature.WriteString(data["amount"])
	rawSignature.WriteString("&extraData=")
	rawSignature.WriteString(data["extraData"])
	rawSignature.WriteString("&ipnUrl=")
	rawSignature.WriteString(data["ipnUrl"])
	rawSignature.WriteString("&orderId=")
	rawSignature.WriteString(data["orderId"])
	rawSignature.WriteString("&orderInfo=")
	rawSignature.WriteString(data["orderInfo"])
	rawSignature.WriteString("&partnerCode=")
	rawSignature.WriteString(data["partnerCode"])
	rawSignature.WriteString("&redirectUrl=")
	rawSignature.WriteString(data["redirectUrl"])
	rawSignature.WriteString("&requestId=")
	rawSignature.WriteString(data["requestId"])
	rawSignature.WriteString("&requestType=")
	rawSignature.WriteString(data["requestType"])

	// Create a new HMAC by defining the hash type and the key (as byte array)
	hmac := hmac.New(sha256.New, []byte(data["secretKey"]))

	// Write Data to it
	hmac.Write(rawSignature.Bytes())
	fmt.Println("Raw signature: " + rawSignature.String())

	// Get result and encode as hexadecimal string
	signature := hex.EncodeToString(hmac.Sum(nil))
	return signature
}

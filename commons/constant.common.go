package commons

var PublicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAkpa+qMXS6O11x7jBGo9W
3yxeHEsAdyDE40UoXhoQf9K6attSIclTZMEGfq6gmJm2BogVJtPkjvri5/j9mBnt
A8qKMzzanSQaBEbr8FyByHnf226dsLt1RbJSMLjCd3UC1n0Yq8KKvfHhvmvVbGcW
fpgfo7iQTVmL0r1eQxzgnSq31EL1yYNMuaZjpHmQuT24Hmxl9W9enRtJyVTUhwKh
tjOSOsR03sMnsckpFT9pn1/V9BE2Kf3rFGqc6JukXkqK6ZW9mtmGLSq3K+JRRq2w
8PVmcbcvTr/adW4EL2yc1qk9Ec4HtiDhtSYd6/ov8xLVkKAQjLVt7Ex3/agRPfPr
NwIDAQAB
-----END PUBLIC KEY-----
`)
var Endpoint = "https://test-payment.momo.vn/v2/gateway/api/create"
var PartnerCode = "MOMOTVT020211201"
var AccessKey = "8aZ8fEbb6xl9hlxu"
var SecretKey = "ZRTaB3s8zKvdvnhWDRpdqPUsDvU6OuLc"
var RequestType = "captureWallet"
var ExtraData = "" //pass empty value or Encode base64 JsonString
var IpnUrl = "https://momo.vn"

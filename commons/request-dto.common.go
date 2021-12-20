package commons

type RequestDTO struct {
	OrderID     string `json:"order_id"`
	OrderInfo   string `json:"order_info"`
	RedirectURL string `json:"redirect_url"`
	Amount      uint64 `json:"amount"`
	PartnerCode string `json:"partner_code"`
	AccessKey   string `json:"access_key"`
	SecretKey   string `json:"secret_key"`
}

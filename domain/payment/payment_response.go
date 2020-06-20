package payment

type CreatePaymentResp struct {
	RedirectUrl string `json:"redirectUrl"`
}

type CompletePaymentResp struct {
	Status string `json:"status"`
}

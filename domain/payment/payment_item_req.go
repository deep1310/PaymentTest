package payment

import "payment/utils/errors"

type PaymentItemReq struct {
	OrderId     int64   `json:"orderId"`
	Amount      float64 `json:"amount"`
	PaymentType string  `json:"paymentType"`
}

type PaymentUpdateReq struct {
	PaymentId int64
	Status    string
}

func (r *PaymentUpdateReq) ValidatePaymentUpdateStatus() *errors.RestErr {

	if r.Status != "SUCCESS" && r.Status != "FAILED" {
		return errors.BadRequestError("invalid request")
	}
	return nil
}

package services

import (
	"payment/domain/payment"
	"payment/utils/errors"
)

var (
	PaymentService paymentServiceInterface = &paymentServiceRepo{}
)

type paymentServiceRepo struct{}

type paymentServiceInterface interface {
	CreatePaymentOrder(*payment.PaymentItemReq) (*payment.CreatePaymentResp, *errors.RestErr)
	UpdatePaymentOrder(*payment.PaymentUpdateReq) *errors.RestErr
	UpdatePaymentTransactionId(string) *errors.RestErr
}

func (s *paymentServiceRepo) CreatePaymentOrder(paymentReq *payment.PaymentItemReq) (*payment.CreatePaymentResp, *errors.RestErr) {

	paymentCreate := &payment.Payment{
		PgName:      "RazorPay",
		OrderId:     paymentReq.OrderId,
		Amount:      paymentReq.Amount,
		PaymentType: paymentReq.PaymentType,
		Status:      "INITIATED",
	}

	if err := paymentCreate.PaymentSave(); err != nil {
		return nil, err
	}

	/*
			Here we now call the pg for first factor authentication
			and it returns the status and pgTxnId
		    and we save the same to our payment db
	*/
	pgTransId := "123"
	if pgIdUpdateErr := s.UpdatePaymentTransactionId(pgTransId); pgIdUpdateErr != nil {
		return nil, pgIdUpdateErr
	}
	paymentResp := &payment.CreatePaymentResp{
		RedirectUrl: "https://razorpay.com",
	}

	return paymentResp, nil
}

func (s *paymentServiceRepo) UpdatePaymentTransactionId(pgTransactionId string) *errors.RestErr {
	pgIdUpdate := &payment.PGIdUpdateReq{
		PgId:   pgTransactionId,
		Status: "SUCCESS",
	}
	if pgTransactionId == "" {
		pgIdUpdate.Status = "FAILED"
	}
	if err := pgIdUpdate.PaymentGatewayUpdate(); err != nil {
		return err
	}
	return nil
}

/*
	this method will be call by the pg as web-hook after the second factor is success
	form here the api gateway complete order will be called
*/

func (s *paymentServiceRepo) UpdatePaymentOrder(req *payment.PaymentUpdateReq) *errors.RestErr {
	paymentUpdate := &payment.Payment{
		Id:     req.PaymentId,
		Status: req.Status,
	}
	paymentUpdate.UpdatePaymentStatus()
	return nil
}

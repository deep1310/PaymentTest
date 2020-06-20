package payment

import (
	"payment/datarepository/mysql/payment_db"
	"payment/utils/errors"
)

func (paymentReq *Payment) PaymentSave() *errors.RestErr {
	db := payment_db.GetSqlConn()
	if result := db.Model(&paymentReq).Create(&paymentReq); result != nil {
		if result.Error != nil {
			return errors.InternalServerError("Not able to create payment item")
		}
	}
	return nil
}

func (pgUpdateReq *PGIdUpdateReq) PaymentGatewayUpdate() *errors.RestErr {
	db := payment_db.GetSqlConn()
	paymentUpdate := Payment{}
	totalRowsUpdated := db.Model(&paymentUpdate).Where("id = ?", pgUpdateReq.PaymentId).Updates(map[string]interface{}{
		"pgTxnId": pgUpdateReq.PgId}).RowsAffected
	if totalRowsUpdated == 0 {
		return errors.InternalServerError("Not able to update payment")
	}
	return nil
}

func (r *Payment) UpdatePaymentStatus() *errors.RestErr {

	if err := r.PaymentGet(); err != nil {
		return err
	}
	db := payment_db.GetSqlConn()

	totalRowsUpdated := db.Model(r).Where("id = ?", r.Id).Updates(map[string]interface{}{
		"finalStatus": r.Status}).RowsAffected
	if totalRowsUpdated == 0 {
		return errors.InternalServerError("Not able to update payment")
	}

	return nil
}

func (paymentReq *Payment) PaymentGet() *errors.RestErr {
	db := payment_db.GetSqlConn()
	err := db.Where("id = ?", paymentReq.Id).Find(&paymentReq).Error
	if err != nil {
		return errors.InternalServerError("Not able to get payment")
	}
	return nil
}

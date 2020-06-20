package payment

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Payment struct {
	Id          int64     `gorm:"primary_key";json:"paymentId"`
	OrderId     int64     `json:"orderId"`
	PaymentType string    `json:"paymentType"`
	PgName      string    `json:"pgName"`
	PgTxnId     string    `json:"pgTxnId"`
	Amount      float64   `json:"amount"`
	Status      string    `json:"status"`
	FinalStatus string    `json:"finalStatus"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

// set Segment table name to be `segments`
func (Payment) TableName() string {
	return "payment_details"
}

func (paymentDetail *Payment) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now().UTC())
	scope.SetColumn("UpdatedAt", time.Now().UTC())
	return nil
}

func (paymentDetail *Payment) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().UTC())
	return nil
}

type PGIdUpdateReq struct {
	PgId      string
	PaymentId int64
	Status    string
}


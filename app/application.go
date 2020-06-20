package app

import (
	"github.com/gin-gonic/gin"
	"payment/controllers/payment"
)

var router = gin.Default()

func Start() {
	router.POST("/CreatePayment", payment.CreatePaymentRequest)
	router.GET("/PgResponse", payment.ProcessPGResp)
	router.Run(":5552")
}

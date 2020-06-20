package payment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment/domain/payment"
	"payment/services"
	"payment/utils/errors"
	"strconv"
	"strings"
)

func CreatePaymentRequest(c *gin.Context) {
	var paymentReq payment.PaymentItemReq
	if err := c.ShouldBindJSON(&paymentReq); err != nil {
		apiReqErr := errors.BadRequestError("invalid request")
		c.JSON(apiReqErr.Code, apiReqErr)
		return
	}

	result, err := services.PaymentService.CreatePaymentOrder(&paymentReq)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func ProcessPGResp(c *gin.Context) {

	paymentId, isPaymentIdExists := c.GetQuery("paymentId")

	if !isPaymentIdExists || paymentId == "" {
		apiReqErr := errors.BadRequestError("invalid request")
		c.JSON(apiReqErr.Code, apiReqErr)
		return
	}

	paymentInt, err := strconv.ParseInt(paymentId, 10, 64)
	if err != nil {
		apiReqErr := errors.BadRequestError("invalid request")
		c.JSON(apiReqErr.Code, apiReqErr)
		return
	}
	status, isStatusExists := c.GetQuery("status")
	status = strings.TrimSpace(strings.ToUpper(status))
	if !isStatusExists || status == "" {
		apiReqErr := errors.BadRequestError("invalid request")
		c.JSON(apiReqErr.Code, apiReqErr)
		return
	}

	req := &payment.PaymentUpdateReq{
		Status:    status,
		PaymentId: paymentInt,
	}
	req.ValidatePaymentUpdateStatus()
	services.PaymentService.UpdatePaymentOrder(req)
	c.JSON(http.StatusOK, "")

}

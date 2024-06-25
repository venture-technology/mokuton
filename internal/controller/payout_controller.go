package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/vtx-payout-microservice-user/internal/service"
	"github.com/venture-technology/vtx-payout-microservice-user/models"
)

type PayoutController struct {
	payoutservice *service.PayoutService
}

func NewPayoutController(payoutservice *service.PayoutService) *PayoutController {
	return &PayoutController{
		payoutservice: payoutservice,
	}
}

func (ct *PayoutController) RegisterRoutes(router *gin.Engine) {

	api := router.Group("vtx-payout-mock-service/api/v1")

	api.GET("/ping", ct.Ping)
	api.POST("/payout", ct.CreatePayout)
	api.GET("/payout/:id", ct.ReadPayout)
	api.DELETE("/payout/:id", ct.DeletePayout)

}

func (ct *PayoutController) Ping(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})

}

func (ct *PayoutController) CreatePayout(c *gin.Context) {

	var input models.Payout

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err := ct.payoutservice.CreatePayout(c, &input)

	if err != nil {
		log.Printf("create payout error: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, input)

}

func (ct *PayoutController) ReadPayout(c *gin.Context) {

	id := c.Param("id")

	payout, err := ct.payoutservice.ReadPayout(c, &id)

	if err != nil {
		log.Printf("error while found payout: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "payout don't found"})
		return
	}

	c.JSON(http.StatusOK, payout)

}

func (ct *PayoutController) DeletePayout(c *gin.Context) {

	id := c.Param("id")

	err := ct.payoutservice.DeletePayout(c, &id)

	if err != nil {
		log.Printf("error while found payout: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "payout don't found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("payout deleted: %v", id),
	})

}

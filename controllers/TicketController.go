package controllers

import (
	"github.com/gin-gonic/gin"
	"sbytes_v3/entities"
	"sbytes_v3/services"
	"time"
)

const ExpirationTicketTime = 1 * time.Minute

type TicketController struct {
	ms *services.MongoService
}

func NewTicketController(ms *services.MongoService) *TicketController {
	return &TicketController{ms: ms}
}

func (r *TicketController) newTicket() *entities.Ticket {
	return &entities.Ticket{
		CreateAt:       time.Now().Unix(),
		ExpirationDate: time.Now().Add(ExpirationTicketTime).Unix(),
	}
}

func (r *TicketController) CreateTicket(ctx *gin.Context) {
	ticket := r.newTicket()

	_id, err := r.ms.InsertTicket(*ticket)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"id": _id})
}

func (r *TicketController) GetTicket(ctx *gin.Context) {
	uriParamTicketId := ctx.Param("id")

	document, err := r.ms.FindTicketAsBsonDocument(uriParamTicketId)

	if isTicketNotFound(ctx, err) {
		return
	}

	if document["expirationDate"].(int64) < time.Now().Unix() {
		ctx.JSON(404, gin.H{"error": "Ticket is expired"})
		return
	}

	ctx.JSON(200, document)
}

func isTicketNotFound(ctx *gin.Context, err error) bool {
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return true
	}
	return false
}

func (r *TicketController) UpdateTicket(ctx *gin.Context) {
	uriParamTicketId := ctx.Param("id")

	var ticket entities.Ticket

	if err := ctx.BindJSON(&ticket); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	document, err := r.ms.FindTicketAsBsonDocument(uriParamTicketId)

	if isTicketNotFound(ctx, err) {
		return
	}

	if document["expirationDate"].(int64) < time.Now().Unix() {
		ctx.JSON(404, gin.H{"error": "Ticket is expired"})
		return
	}

	if _, err := r.ms.UpdateTicket(uriParamTicketId, ticket); err != nil {
		ctx.JSON(400, gin.H{"error": err})
		return
	}

	ctx.JSON(200, gin.H{"id": uriParamTicketId})
}

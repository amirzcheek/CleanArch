package http

import (
	useCase "architecture_go/services/contact/internal/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type Delivery struct {
	ucContact      useCase.Contact
	ucGroup        useCase.Group
	ucContactGroup useCase.ContactGroup
	Router         *gin.Engine
}

func New(ucContact useCase.Contact, ucGroup useCase.Group, ucContactGroup useCase.ContactGroup) *Delivery {
	var del = &Delivery{
		ucContact:      ucContact,
		ucGroup:        ucGroup,
		ucContactGroup: ucContactGroup,
	}
	del.Router = del.initRouter()
	return del
}

func (del *Delivery) Run() error {
	return del.Router.Run(fmt.Sprintf("%d", uint16(viper.GetUint("HTTP_PORT"))))
}

func checkAuthentication(c *gin.Context) {
	c.Next()
}

func (del *Delivery) handleGetContacts(c *gin.Context) {
}

func (del *Delivery) handleCheckExistingContact(c *gin.Context) {
	name := c.Param("name")
	if err := del.ucContact.CheckExistingContact(name); err != nil {
		log.Printf("Error checking existing contact: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (del *Delivery) handleCreateContact(c *gin.Context) {
	var req CreateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	traceID := uuid.New().String()
	log.Printf("[%s] Started tracing for creating contact", traceID)

	createdContact, err := del.ucContact.CreateContact(traceID, req.Name)
	if err != nil {
		log.Printf("[%s] Error creating contact: %v", traceID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Отправка ответа
	c.JSON(http.StatusCreated, gin.H{"contact": createdContact})
}

func (del *Delivery) handleUpdateContact(c *gin.Context) {
}

func (del *Delivery) handleDeleteContact(c *gin.Context) {
}

func (del *Delivery) handleGetGroups(c *gin.Context) {
}

func (del *Delivery) handleCreateGroup(c *gin.Context) {
}

func (del *Delivery) handleUpdateGroup(c *gin.Context) {
}

func (del *Delivery) handleDeleteGroup(c *gin.Context) {
}

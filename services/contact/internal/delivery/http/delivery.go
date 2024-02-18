package http

import (
	"architecture_go/services/contact/internal/useCase"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Delivery struct {
	ucContact useCase.Contact
	ucGroup   useCase.Group
	router    *gin.Engine
}

func New(ucContact useCase.Contact, ucGroup useCase.Group) *Delivery {
	var del = &Delivery{
		ucContact: ucContact,
		ucGroup:   ucGroup,
	}
	del.router = del.initRouter()
	return del
}

func (del *Delivery) Run() error {
	return del.router.Run(fmt.Sprintf("%d", uint16(viper.GetUint("HTTP_PORT"))))
}

func checkAuthentication(c *gin.Context) {
	c.Next()
}

func (del *Delivery) handleGetContacts(c *gin.Context) {
}

func (del *Delivery) handleCreateContact(c *gin.Context) {
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

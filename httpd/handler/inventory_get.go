package handler

import (
	"backend/platform/inventory"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InventoryGet(inv inventory.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := inv.GetAll()
		c.JSON(http.StatusOK, results)

	}
}

//create items within a container

package middleware

import (
	"github.com/gin-gonic/gin"
	"it-planet-task/internal/app/model/entity"
	"net/http"
)

func AdminRequired(c *gin.Context) {
	authorizedAccountAny, _ := c.Get("account")
	authorizedAccount := authorizedAccountAny.(*entity.Account)
	if authorizedAccount.Role != entity.AdminRole {
		c.AbortWithStatusJSON(http.StatusForbidden, "Only admin can access this endpoint")
		return
	}

	c.Next()
}

func AdminOrChipperRequired(c *gin.Context) {
	authorizedAccountAny, _ := c.Get("account")
	authorizedAccount := authorizedAccountAny.(*entity.Account)
	if authorizedAccount.Role != entity.AdminRole && authorizedAccount.Role != entity.ChipperRole {
		c.AbortWithStatusJSON(http.StatusForbidden, "Only admin or chipper can access this endpoint")
		return
	}

	c.Next()
}

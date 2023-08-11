package v1

import (
	"github.com/gin-gonic/gin"
	"server/app/models/link"
	"server/pkg/response"
)

type LinksController struct {
	BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	response.Data(c, link.AllCached())
}

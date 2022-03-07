package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/link"
	"gohub/pkg/response"
)

type LinksController struct {
	BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	response.Data(c, link.AllCached())
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type rootController struct{}

func NewRootController() rootController {
	return rootController{}
}

func (r *rootController) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "root/index.tmpl", gin.H{})
}

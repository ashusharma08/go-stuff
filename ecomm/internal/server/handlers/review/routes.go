package review

import "github.com/gin-gonic/gin"

type reviewRouter struct {
}

func NewreviewRouter() *reviewRouter {
	return &reviewRouter{}
}

func (p *reviewRouter) RegisterRoutes(r *gin.RouterGroup) {
	// register routes and handlers
}

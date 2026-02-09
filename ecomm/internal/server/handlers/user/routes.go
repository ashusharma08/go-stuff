package user

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func NewUserRouter() *UserRouter {
	return &UserRouter{}
}

func (p *UserRouter) RegisterRoutes(r *gin.RouterGroup) {
	// register routes and handlers
}

package shoppingcart

import "github.com/gin-gonic/gin"

type shoppingcartRouter struct {
}

func NewshoppingcartRouter() *shoppingcartRouter {
	return &shoppingcartRouter{}
}

func (p *shoppingcartRouter) RegisterRoutes(r *gin.RouterGroup) {
	// register routes and handlers
}

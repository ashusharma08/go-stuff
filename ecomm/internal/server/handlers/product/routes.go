package product

import "github.com/gin-gonic/gin"

type ProductRouter struct {
}

func NewProductRouter() *ProductRouter {
	return &ProductRouter{}
}

func (p *ProductRouter) RegisterRoutes(r *gin.RouterGroup) {
	//register routes and handlers
}

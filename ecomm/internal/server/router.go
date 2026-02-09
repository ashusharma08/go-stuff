package server

import (
	"github.com/esoptra/go-prac/ecomm/internal/server/handlers/product"
	"github.com/esoptra/go-prac/ecomm/internal/server/handlers/review"
	"github.com/esoptra/go-prac/ecomm/internal/server/handlers/shoppingcart"
	user "github.com/esoptra/go-prac/ecomm/internal/server/handlers/user"

	"github.com/gin-gonic/gin"
)

type Router interface {
	RegisterRoutes(r *gin.RouterGroup)
}

func GetRoutes() *gin.Engine {
	r := gin.Default()
	gp := r.Group("/api/v1")
	routers := make([]Router, 0)
	routers = append(routers, shoppingcart.NewshoppingcartRouter(), user.NewUserRouter(), product.NewProductRouter(), review.NewreviewRouter())
	for _, rt := range routers {
		rt.RegisterRoutes(gp)
	}
	return r
}

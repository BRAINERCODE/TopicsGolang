package main

import (
	LG "Topic/Businesslogic"

	"github.com/gin-gonic/gin"
)

func main() {
	var i LG.ImplementBussines
	c := gin.Default()

	c.GET("/productos", i.GetAll)             //Done
	c.GET("/productos/:id", i.GetProductById) //Done
	c.POST("/productos", i.SaveProduct)       //Done
	c.PUT("/productos/:id", i.Update)         //Done
	c.DELETE("productos/:id", i.Delete)       //Done
	c.Run(":8090")

}

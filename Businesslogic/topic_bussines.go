package bussineslogic

import (
	data "Topic/Data"
	"Topic/Models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ImplementBussines struct {
	Data data.Implement
}

var p Models.Topic
var Array []Models.Topic

func (i ImplementBussines) GetAll(c *gin.Context) {

	Array, err := i.Data.GetAllTopics()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {

		c.JSON(http.StatusOK, gin.H{
			"values": Array,
		})
	}

}

func (i ImplementBussines) GetProductById(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	GetById, err := i.Data.GetTopicById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {

		c.JSON(http.StatusOK, gin.H{
			"values": GetById,
		})
	}
}
func (i ImplementBussines) SaveProduct(c *gin.Context) {

	err := c.ShouldBindJSON(&p)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	i.Data.PostTopic(p)
}

func (i ImplementBussines) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	i.Data.UpdateTopic(p, int(id))

}

func (i ImplementBussines) Delete(c *gin.Context) {

	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	err := i.Data.DeleteTopic(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

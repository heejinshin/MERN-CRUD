package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackbalageru/MERN-CRUD/go-server/pkg/api"
)

func Router(apis *api.APIs) *gin.Engine {
	r := gin.Default()

	crudRouter := r.Group("/api/cruds")

	crudRouter.GET("/", apis.GetBoardList)
	crudRouter.POST("/", apis.CreateBoard)
	crudRouter.GET("/:id", apis.GetBoardByID)
	crudRouter.PATCH("/:id", apis.UpdateBoard)
	crudRouter.DELETE("/:id", apis.DeleteBoardByID)

	return r
}

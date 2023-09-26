package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/railanderreis/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/aluno", controllers.CriaNovoAluno)
	r.Run()
}

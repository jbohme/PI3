package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jbohme/crud/internal/app/http/controller"
	"github.com/jbohme/crud/internal/app/model"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {
	//r.Group("/user")
	//{
	r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", model.VerifyTokenMiddleware, userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", model.VerifyTokenMiddleware, userController.DeleteUser)

	r.POST("/login", userController.LoginUser)

	// Rota WebSocket para entrar em uma sala aleatória
	r.GET("/JoinRandomRoom", userController.JoinRandomRoom)

	//}
}

package routes

import (
	"jwt_golang/controllers"
	"jwt_golang/models"
)

var authController = controllers.NewAuthController(models.GetSession())

func init() {

	BASE_ROUTER.HandleFunc("/auth", authController.Authenticate).Methods("POST")

}

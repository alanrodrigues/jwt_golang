package routes

import (
	"jwt_golang/controllers"
	"jwt_golang/models"

	"github.com/urfave/negroni"
)

var offerController = controllers.NewOfferController(models.GetSession())

func init() {

	authController.InitKeys()


	BASE_ROUTER.HandleFunc("/offers", offerController.GetOffers).Methods("GET")
	SECURED_ROUTER.HandleFunc("/offers", offerController.SaveOffer).Methods("POST")

	BASE_ROUTER.PathPrefix("/secured").Handler(negroni.New(
		negroni.HandlerFunc(authController.ValidateTokenMiddleware),
		negroni.Wrap(SECURED_ROUTER),
	))
	
}
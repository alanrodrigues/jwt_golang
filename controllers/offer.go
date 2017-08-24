package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"jwt_golang/models"
	"jwt_golang/helpers/response"

	mgo "gopkg.in/mgo.v2"

)

type (
	OfferController struct {
		session *mgo.Session
	}
)

func NewOfferController(s *mgo.Session) *OfferController {
	return &OfferController{s}
}

func (oc OfferController) GetOffers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var offers [] models.Offer

	if err := oc.session.DB("jwt_golang_db").C("offers").Find(nil).All(&offers); err != nil {
        w.WriteHeader(500)
        return
	}

	if offers == nil {
		response.EmptyJsonArrayResponse(w)
		return
	}

	response.JsonResponse(offers, w);
}

func (oc OfferController) SaveOffer(w http.ResponseWriter, r *http.Request) {


	decoder := json.NewDecoder(r.Body)

    var offer models.Offer
    err := decoder.Decode(&offer)

    if err != nil {
        w.WriteHeader(500)
        fmt.Println(err)
        return
    }

    defer r.Body.Close()

    if err := oc.session.DB("jwt_golang_db").C("offers").Insert(offer); err != nil {
        w.WriteHeader(500)
        fmt.Println(err)
        return
	}

	response.JsonResponse(offer, w)
}


package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"jwt_golang/models"
	"jwt_golang/helpers/response"
	"log"

	"bytes"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/dgrijalva/jwt-go"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

)

var signingKey, verificationKey []byte

type UserCredentials struct {
	Username	string  `json:"username"`
	Password	string	`json:"password"`
}

type Token struct {
	Token 	string    `json:"token"`
}

type (
	AuthController struct {
		session *mgo.Session
	}
)

func NewAuthController(s *mgo.Session) *AuthController {
	return &AuthController{s}
}

func (ac AuthController) Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var user UserCredentials

	//decode request into UserCredentials struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Error in request")
		return
	}

	//validate user credentials
	var registeredUser models.User

	if err := ac.session.DB("jwt_golang_db").C("users").Find(bson.M{"username":user.Username, "password":user.Password}).One(&registeredUser); err != nil {
        w.WriteHeader(http.StatusForbidden)
		fmt.Println("Error logging in")
		fmt.Fprint(w, "Invalid credentials")
        return
	}

	token := jwt.New(jwt.SigningMethodHS256)
    claims := make(jwt.MapClaims)
    claims["iss"] = "admin"
    claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
    claims["iat"] = time.Now().Unix()
    claims["jti"] = "1" // should be user ID?
    token.Claims = claims

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		log.Printf("Error signing token: %v\n", err)
	}

	//create a token instance using the token string
	resp := Token{tokenString}
	response.JsonResponse(resp, w)

}

func (ac AuthController) InitKeys() {

	var (
		err         error
		privKey     *rsa.PrivateKey
		pubKey      *rsa.PublicKey
		pubKeyBytes []byte
	)

	privKey, err = rsa.GenerateKey(cryptorand.Reader, 2048)
	if err != nil {
		log.Fatal("Error generating private key")
	}
	pubKey = &privKey.PublicKey //hmm, this is stdlib manner...

	// Create signingKey from privKey
	// prepare PEM block
	var privPEMBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey), // serialize private key bytes
	}
	// serialize pem
	privKeyPEMBuffer := new(bytes.Buffer)
	pem.Encode(privKeyPEMBuffer, privPEMBlock)
	//done
	signingKey = privKeyPEMBuffer.Bytes()

	// create verificationKey from pubKey. Also in PEM-format
	pubKeyBytes, err = x509.MarshalPKIXPublicKey(pubKey) //serialize key bytes
	if err != nil {
		// heh, fatality
		log.Fatal("Error marshalling public key")
	}

	var pubPEMBlock = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	}
	// serialize pem
	pubKeyPEMBuffer := new(bytes.Buffer)
	pem.Encode(pubKeyPEMBuffer, pubPEMBlock)
	// done
	verificationKey = pubKeyPEMBuffer.Bytes()

}

func (ac AuthController) ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

    var responseToken = r.Header.Get("Authorization")
	
	token, err := jwt.Parse(responseToken, func(token *jwt.Token) (interface{}, error){
		return signingKey, nil
	})

	if err == nil {

		if token.Valid{
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorised access to this resource")
	}

}
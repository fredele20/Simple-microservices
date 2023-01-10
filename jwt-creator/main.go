package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func GetJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Victor"
	claims["aud"] = "billing.jwt.io"
	claims["iss"] = "jwtgo.to"
	claims["exp"] = time.Now().Add(time.Minute*1).Unix()

	tokenString, err := token.SignedString(MySigningKey)
	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func Index(w http.ResponseWriter, r *http.Request) {
	validToken, err := GetJWT()
	fmt.Println(validToken)
	if err != nil {
		fmt.Println("failed to generate token")
	}
	fmt.Fprintf(w, string(validToken))
}

func handleRequests() {
	http.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequests()
}
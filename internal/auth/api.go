package auth

import (
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	user2 "github.com/vediagames/gamalytics/internal/user"
	"github.com/vediagames/gamalytics/pkg/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

var secretJWT = []byte(helpers.Config().JWT.Secret)

var connectUri =helpers.Config().Mongodb.ConnectionStrings.User
var client *mongo.Client
var database *mongo.Database

func Init(port string){
	client, err := mongo.NewClient(options.Client().ApplyURI(connectUri))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	helpers.HandleError(err)
	database = client.Database("user")

	r := mux.NewRouter()
	r.HandleFunc("/login", userLogIn).Methods(http.MethodPost)
	r.HandleFunc("/validate",validateCredentials).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":" + port, r))

}

func userLogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user user2.User
	var dbUser user2.User

	json.NewDecoder(r.Body).Decode(&user)
	collection := database.Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil {
		log.Println(passErr)
		w.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}

	jwtToken, err := helpers.GenerateJWT(secretJWT)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	w.Write([]byte(`{"token":"` + jwtToken + `"}`))

}

func validateCredentials(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var rToken string
	json.NewDecoder(r.Body).Decode(rToken)

	token, err := jwt.Parse(rToken, func(jwtToken *jwt.Token)(interface{}, error){
		return secretJWT, nil
	})

	if err == nil && token.Valid {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("What are you doing step bro"))
	}
}
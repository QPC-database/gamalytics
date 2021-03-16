package user

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vediagames/gamalytics/pkg/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)



type User struct{
	FirstName string `json:"firstname" bson:"firstname"`
	LastName string `json:"lastname" bson:"lastname"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

var connectUri = helpers.Config().Mongodb.ConnectionStrings.User
var client *mongo.Client
var database *mongo.Database

func Init(port string){
	client, err := mongo.NewClient(options.Client().ApplyURI(connectUri))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	helpers.HandleError(err)
	database = client.Database("user")

	r := mux.NewRouter()
	r.HandleFunc("/new", newUser).Methods(http.MethodPost)
	r.HandleFunc("/{userid}", getUser).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":"+ port, r))

}

func newUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	user.Password = helpers.GetHash([]byte(user.Password))

	collection := database.Collection("user")

	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	result,_ := collection.InsertOne(ctx,user)

	json.NewEncoder(w).Encode(result)
}

func getUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	var user User

	collection := database.Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := collection.FindOne(ctx, bson.M{"email": vars["userid"]}).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	userString, err := json.Marshal(user)
	helpers.HandleError(err)
	w.WriteHeader(http.StatusOK)
	w.Write(userString)
}


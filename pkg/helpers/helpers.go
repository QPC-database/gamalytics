package helpers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"os"
)

type ConfigFile struct  {
	Mongodb struct {
		ConnectionStrings struct {
			User string `json:"user"`
		} `json:"connectionStrings"`
	} `json:"mongodb"`

	JWT struct {
		Secret string `json:"secret"`
	} `json:"jwt"`

}

// HandleError - Basic error handler
func HandleError(err error) {
	if err != nil{
		log.Fatal(err)
	}
}

func GetHash(pwd []byte) (string) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	HandleError(err)
	return string(hash)
}


func GenerateJWT(secretKey interface{})(string,error){
	// Create the token
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}

func Config() (*ConfigFile){
	jsonFile, err := os.Open("configs/default.json")
	HandleError(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var c ConfigFile
	err = json.Unmarshal([]byte(byteValue), &c)
	HandleError(err)

	return &c
}
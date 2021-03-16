package main

import (
	"fmt"
	"github.com/vediagames/gamalytics/internal/auth"
)

func main(){
	port := "10000"
	fmt.Println("Starting auth service on port: " + port)
	auth.Init(port)
}
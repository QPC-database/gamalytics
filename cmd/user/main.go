package main

import (
	"fmt"
	"github.com/vediagames/gamalytics/internal/user"
)

func main()  {
	port := "10001"
	fmt.Println("Starting user service on port: " + port)
	user.Init(port)
}

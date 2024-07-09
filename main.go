package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vikky225/mongo-golang/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	router := httprouter.New()
	uc := controllers.NewUserController(getSession())
	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe(":8090", router)
}

func getSession() *mgo.Session {

	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	return session
}

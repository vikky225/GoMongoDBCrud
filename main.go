package main

import (
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
	http.ListenAndServe(":8080", router)
}

func getSession() *mgo.Session {

	session, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}
	return session
}

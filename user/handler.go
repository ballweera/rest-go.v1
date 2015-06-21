package user

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"gopkg.in/mgo.v2"
)

func Add(w rest.ResponseWriter, r *rest.Request) {
	user := User{}
	err := r.DecodeJsonPayload(&user)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("tidtor").C("users")
	err = c.Insert(&user)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteJson(&user)
}

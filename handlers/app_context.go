package handlers

import "gopkg.in/mgo.v2"

//import (
//	"encoding/json"
//	"github.com/julienschmidt/httprouter"
//	"github.com/gorilla/context"
//	"net/http"
//	"gopkg.in/mgo.v2"
//	"kommunalka-server/repo"
//)

type AppContext struct {
	DB *mgo.Database
}

//func (c *appContext) authHandler(next http.Handler) http.Handler {
//	fn := func(w http.ResponseWriter, r *http.Request) {
//		authToken := r.Header.Get("Authorization")
//		log.Println(authToken)
//		user, err := getUser(c.db, authToken)
//
//		if err != nil {
//			http.Error(w, http.StatusText(401), 401)
//			return
//		}
//
//		context.Set(r, "user", user)
//		next.ServeHTTP(w, r)
//	}
//
//	return http.HandlerFunc(fn)
//}
//
//func getUser(db *mgo.Database, authToken string) (string, error) {
//	return "user1", nil
//}

//func (c *appContext) billsHandler(w http.ResponseWriter, r *http.Request) {
//	params := context.Get(r, "params").(httprouter.Params)
//	year := params.ByName("year")
//	repo := repo.BillRepo{c.db.C("bills")}
//	bills, err := repo.All(year)
//	if err != nil {
//		panic(err)
//	}
//
//	w.Header().Set("Content-Type", "application/vnd.api+json")
//	json.NewEncoder(w).Encode(bills)
//}

//func (c *AppContext) BillsHandler(w http.ResponseWriter, r *http.Request) {
//	params := context.Get(r, "params").(httprouter.Params)
//	year := params.ByName("year")
//	repo := repo.BillRepo{c.DB.C("bills")}
//	billsC, err := repo.All(year)
//	if err != nil {
//		panic(err)
//	}
//
//	w.Header().Set("Content-Type", "application/vnd.api+json")
//	json.NewEncoder(w).Encode(billsC.Data)
//}

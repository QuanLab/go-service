package main

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"

	"github.com/QuanLab/go-service/config"
	"github.com/QuanLab/go-service/service"
	"github.com/QuanLab/go-service/model"
)

func getListId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var accessToken = r.Header.Get("access_token")
		user := service.ValidateToken(accessToken)

		if user.Role == model.GUEST {
			w.Write([]byte("{'message' : 'Current user is not authenticated.'}"))
		}

		if user.Role == model.BASIC {

		}

		if user.Role == model.ADVANCE {

		}

		if user.Role == model.PRO {

		}

		if (user.Role == model.ADMIN) {
			result, err := json.Marshal(user)
			if err != nil {
				log.Panic(err)
			}
			w.Write(result)
		}
	default:
		w.Write([]byte("{'message', 'Method is not allowed here'}"))
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		username := r.FormValue("username")
		password := r.FormValue("password")
		user := service.Login(username, password)
		result, err := json.Marshal(user)
		if err != nil {
			log.Fatal("{'error' : 100, 'message' : 'Fail to login'}")
		}
		w.Write([]byte(result))
	default:
		w.Write([]byte("{'message', 'Method is not allowed here'}"))
	}
}

func main() {
	var conf = config.Get().Server
	http.HandleFunc(conf.BaseContextPath+"getListId", getListId)
	http.HandleFunc(conf.BaseContextPath+"login", login)
	log.Printf("Listen and serve at %d", conf.Port)
	var err = http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil)
	if err != nil {
		panic(err)
	}
}

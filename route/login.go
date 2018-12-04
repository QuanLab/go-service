package route

import (
	"net/http"
	"github.com/QuanLab/go-service/service"
	"encoding/json"
	"log"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

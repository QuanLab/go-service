package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/QuanLab/go-service/config"
	"github.com/QuanLab/go-service/route"
)

func main() {
	var conf = config.Get().Server
	http.HandleFunc(conf.BaseContextPath+"register", route.Register)
	http.HandleFunc(conf.BaseContextPath+"verify", route.VerifyRegister)
	http.HandleFunc(conf.BaseContextPath+"login", route.Login)
	http.HandleFunc(conf.BaseContextPath+"getListId", route.GetListId)
	log.Printf("Listen and serve at %d", conf.Port)
	var err = http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil)
	if err != nil {
		panic(err)
	}
}

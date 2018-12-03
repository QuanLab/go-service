package main

import (
	"fmt"
	"github.com/QuanLab/go-service/config"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

const (
	baseContextPath = "/api/v1/";
)

func main() {
	fmt.Println(config.Get().MysqlPort)
	http.HandleFunc(baseContextPath  + "getListId", sayHello)
	log.Println("Listen and serve at 8080")
	var err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

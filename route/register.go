package route

import (
	"github.com/QuanLab/go-service/service"
	"net/http"
	"github.com/QuanLab/go-service/model"
	"fmt"
	"github.com/QuanLab/go-service/crypt"
	"github.com/QuanLab/go-service/config"
	"time"
	"strings"
	"strconv"
	"log"
)

func Register(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case http.MethodPost:
		var username= r.Header.Get("username")
		var password= r.Header.Get("password")
		var fullName= r.Header.Get("full_name")
		var email= r.Header.Get("email")
		var phone= r.Header.Get("phone")
		var user = model.User{
			Username: username,
			Password: password,
			FullName: fullName,
			Email: email,
			PhoneNumber: phone,
			Role:model.BASIC,
		}
		if !model.CheckUserExists(email) {
			id := model.InsertOne(user)
			user.ID = id
			go sendVerificationEmail(user)
			w.Write([]byte("{'message' : 'Register successfully, please check email to fully active account'}"))
		}
	}
}


func VerifyRegister(w http.ResponseWriter, r *http.Request)  {
	var tokenVerification = r.URL.Query().Get("token")
	plaintext, err := crypt.Decrypt(tokenVerification)
	if err != nil {
		panic(err)
	}

	var splits = strings.Split(plaintext, "|")
	if len(splits) == 2 {
		id, _ := strconv.ParseInt(splits[0], 10, 64)
		timeSend, _ := strconv.ParseInt(splits[0], 10, 64)
		if (time.Now().UnixNano() - timeSend) > time.Duration(time.Hour * 48).Nanoseconds() {
			log.Printf("Active user with id %d ", id)
			model.UpdateActiveStatus(model.User{ID:id})
		}

	}
	w.Write([]byte("{'message' : 'Your account is activated. You can login right now.'}"))
}

// send email verification to user
func sendVerificationEmail(user model.User) {
	subject := "Edge Tech Vietnam - Email verification"
	log.Println(user.Email)
	request := service.NewRequest([]string{user.Email}, subject)

	var info = fmt.Sprintf("%d|%d", user.ID, time.Now().UnixNano())
	var verificationToken = crypt.Encrypt([]byte(info))
	request.Send("templates/email.html", map[string]string{
		"username":user.FullName,
		"url" : config.Get().MailServer.Domain + config.Get().Server.BaseContextPath + "verify?token=" + verificationToken,
	})
}
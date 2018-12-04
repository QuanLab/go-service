package service

import (
	"fmt"
	"strings"
	"strconv"

	"github.com/QuanLab/go-service/model"
	"github.com/QuanLab/go-service/crypt"
	"github.com/QuanLab/go-service/config"
)

func Login(username string, hashPassword string) model.User {
	var user = model.GetUser(username)
	var passwordHasing = crypt.CreateHash(hashPassword + config.Get().Server.Salt)

	if user.Password == passwordHasing {
		var accessTokenStr = fmt.Sprintf("%d", user.ID) + "|" + user.Username + "|" +  user.Role
		var accessToken = crypt.Encrypt([]byte(accessTokenStr))
		return model.User{Username: user.Username, FullName:user.FullName, AccessToken: accessToken}
	}
	return model.User{ID:int64(-1)};
}

func ValidateToken(accessToken string) model.User {
	tokenInfo, err := crypt.Decrypt(accessToken)
	if err != nil {
		return model.User{ID: int64(-1)}
	}
	var splits []string = strings.Split(tokenInfo, "|")
	if len(splits) >= 3 {
		id, _ := strconv.ParseInt(splits[0], 10, 64)
		var username = splits[1]
		var role = splits[2]
		return model.User{ID: id, Username:username, Role:role}
	}
	return model.User{Role: model.GUEST}
}
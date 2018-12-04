package model

import (
	"github.com/QuanLab/go-service/database/mysql"
	"log"
)

type User struct {
	ID          int64    `json:"id,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Role        string `json:"role,omitempty"`
	IsActive    int    `json:"is_active,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}

func GetUser(username string) User {
	query := "SELECT ID, FULL_NAME, USERNAME, PASSWORD, EMAIL, PHONE_NUMBER, `ROLE`, IS_ACTIVE FROM USER_INFO WHERE USERNAME=?;"
	rows, err := mysql.DB.Query(query, username)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var user User
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.FullName, &user.Username, &user.Password, &user.Email, &user.PhoneNumber, &user.Role, &user.IsActive)
		if err != nil {
			log.Println(err)
			return user
		}
	}
	return user
}

func CheckUserExists(email string) bool {
	query := "SELECT ID FROM USER_INFO WHERE EMAIL = ?"
	rows, err := mysql.DB.Query(query,email)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	if rows.Next() {
		return true
	}
	return false
}

func InsertOne(user User) int64 {
	query := "INSERT INTO USER_INFO (FULL_NAME, USERNAME, PASSWORD, EMAIL, PHONE_NUMBER, `ROLE`, POINT, IS_ACTIVE) " +
		"VALUES(?, ?, ?, ?, ?, ?, 0, 0);"
	result, err := mysql.DB.Exec(query, user.FullName, user.Username, user.Password, user.Email, user.PhoneNumber, user.Role)
	if err != nil {
		log.Println(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return id
}

func UpdateActiveStatus(user User) {
	query := "UPDATE USER_INFO SET IS_ACTIVE = 1 WHERE ID =?"
	_, err := mysql.DB.Exec(query, user.ID)
	if err != nil {
		log.Println(err)
	}
}

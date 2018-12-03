package model

import (
	"github.com/QuanLab/go-service/database/mysql"
	"log"
)


type User struct {
	ID          string `json:"id,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	Username    int64  `json:"username,omitempty"`
	Password    int64  `json:"password,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	IsActive    int    `json:"is_active,omitempty"`
}

func GetUser() User {
	query := "SELECT id, full_name, username, password, phone_number, is_active FROM user_info LIMIT 1;"
	rows, err := mysql.DB.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var user User
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			log.Println(err)
			return user
		}
	}
	return user
}

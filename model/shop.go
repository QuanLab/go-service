package model

import (
	"github.com/QuanLab/go-service/database/mysql"
	"log"
)

type Shop struct {
	Data []string `json:"data,omitempty"`
}

func GetListShopToFollow(userId int) []string {
	var query = "SELECT SHOP_ID FROM USER_SHOP WHERE ID NOT IN ( " +
		"SELECT DISTINCT A.SHOP_ID FROM FOLLOW AS A " +
		"LEFT JOIN (" +
		"SELECT ID AS SHOP_ID FROM USER_SHOP WHERE USER_ID =? " +
		") AS B " +
		"ON A.FOLLOWER_ID = B.SHOP_ID);";

	rows, err := mysql.DB.Query(query, userId)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var result = make([]string, 0)
	for rows.Next() {
		var shopId string
		err := rows.Scan(&shopId)
		if err != nil {
			log.Println(err)
		}
		result = append(result, shopId)
	}
	return result
}

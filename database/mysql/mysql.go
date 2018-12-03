package mysql

import (
	"database/sql"
	"fmt"
	"github.com/QuanLab/go-service/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MySQLInfo struct {
	Username  string
	Password  string
	Name      string
	Hostname  string
	Port      int
	Parameter string
}

var (
	DB  *sql.DB
	DB2 *sql.DB
)

func DataSourceName() string {
	// Example: root:@tcp(localhost:3306)/test
	return config.MysqlUserName +
		":" +
		config.MysqlPassword +
		"@tcp(" +
		config.MysqlHost +
		":" +
		fmt.Sprintf("%d", config.MysqlPort) +
		")/" +
		config.DatabaseName + config.Parameter
}


//open connection to MySQL database, it 's self contains pool init
func Connect(d MySQLInfo) {
	var err error
	DB, err = sql.Open("mysql", DataSourceName())
	if err != nil {
		log.Printf("Cannot connect to MySQL server %s", err)
	}
}

package db

import (
	"database/sql"
	"fmt"

	"github.com/Jbfaneto/api-go/configs"
	_ "github.com/lib/pq"
)

// function to open connection with database
func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	//it returns the connection to the pool or the error
	return conn, err
}

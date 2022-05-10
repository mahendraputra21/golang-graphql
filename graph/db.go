package graph

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-pg/pg/v10"
)

//function connect to postgreSQL Hasura cloud
func Connect() *pg.DB {

	connStr := os.Getenv("DB_URL")
	opt, err := pg.ParseURL(connStr)

	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	if _, DBStatus := db.Exec("SELECT 1"); DBStatus != nil {
		panic("PostgreSQL Hasura is down")
	}

	return db
}

//function connect to postgreSQL DB Local
func ConnectLocal() (con *pg.DB) {

	dbName := os.Getenv("dbName")
	user := os.Getenv("user")
	pass := os.Getenv("pass")
	host := os.Getenv("host")
	port := os.Getenv("port")
	poolSize, err := strconv.Atoi(os.Getenv("poolSize"))

	if err != nil {
		log.Fatalf("Error convert poolsize from string to int %v", err)
	}

	address := fmt.Sprintf("%s:%s", host, port)
	options := &pg.Options{
		User:     user,
		Password: pass,
		Addr:     address,
		Database: dbName,
		PoolSize: poolSize,
	}

	con = pg.Connect(options)
	if con == nil {
		panic("PostgreSQL local is down")
	}

	return
}

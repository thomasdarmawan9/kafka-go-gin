package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func InitializeDB() {

	dbdriver := os.Getenv("DBDRIVER")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	database := os.Getenv("DATABASE")
	PORT := os.Getenv("PORT")

	DBURL := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", username, password, host, PORT, database)

	db, err = sql.Open(dbdriver, DBURL)

	if err != nil {
		log.Fatal("Error connecting to database:", err.Error())
	}

	fmt.Println("Successfully connected to database")

	err = db.Ping()

	if err != nil {
		log.Fatal("Error Ping Database:", err.Error())
	}

	createTable := `CREATE TABLE IF NOT EXISTS orders(
		id  SERIAL PRIMARY KEY,
		foodId int NOT NULL,
		userId int NOT NULL,
		amount smallint NOT NULL,
		total_price numeric NOT NULL,
		status varchar(20) DEFAULT 'UNPAID' NOT NULL,
		created_at timestamptz DEFAULT now()
	)`
	_, err = db.Exec(createTable)

	if err != nil {
		log.Fatal("error creating order table:", err.Error())
	}

	createTable = `CREATE TABLE IF NOT EXISTS foods(
		id  int PRIMARY KEY,
		name varchar(200) NOT NULL,
		price numeric NOT NULL,
		stock smallint NOT NULL,
		created_at timestamptz DEFAULT now()
	)`

	_, err = db.Exec(createTable)

	if err != nil {
		log.Fatal("error creating order table:", err.Error())
	}

}

func GetDB() *sql.DB {
	return db
}

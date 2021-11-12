package mysql

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type Mysql struct {
	DB *sql.DB
}

func NewMySql() *Mysql {
	s, err := dbConnection()
	if err != nil {
		panic(err.Error())
	}
	return s
}
func dbConnection() (*Mysql, error) {
	var db *gorm.DB
	err := godotenv.Load(".env")
	dbUri := os.Getenv("db_uri")
	dbPass := os.Getenv("db_pass")
	dbUser := os.Getenv("db_user")
	log.Print(dbUri)
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbUri + ":3308)/news?charset=utf8mb4&parseTime=True&loc=Local"
	log.Print(dsn)
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	db.LogMode(true)

	return &Mysql{
		DB: db.DB(),
	}, err
}

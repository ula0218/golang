package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
    // 加載 .env 文件中的環境變數
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
    initDB()
    defer db.Close()
}

func initDB() {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbName := os.Getenv("DB_NAME")

    dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

    var err error
    db, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatalf("Error opening database connection: %v", err)
    }
    fmt.Println("DB結構已建立")

    err = db.Ping()
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    fmt.Println("DB連線成功")
    DBctrate :=`
    CREATE TABLE member
    (
        id INT NOT NULL UNIQUE,
        name VARCHAR(20)
    );`
    _ , err = db.Exec(DBctrate) 
    if err != nil{
        log.Fatalf("Error create table : %v", err)
    }
    fmt.Print("資料表已建立")
}

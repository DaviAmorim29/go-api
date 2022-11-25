package main

import (
	"database/sql"

	"github.com/daviamorim29/cookieapi/internal/infra/database"
	"github.com/daviamorim29/cookieapi/internal/infra/webserver/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.New()
	db, err := sql.Open("sqlite3", "./cookieapi.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cookies (id VARCHAR(255), abck TEXT, user_agent TEXT, proxy TEXT, created_at date, updated_at date)")
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	cookieDb := database.NewCookieDB(db)
	handlers.NewCookieHandler(router, cookieDb)
	router.Run(":3002")
}

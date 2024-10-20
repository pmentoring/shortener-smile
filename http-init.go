package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shortener-smile/database"
	"shortener-smile/internal/common"
	"shortener-smile/internal/routing"
)

func main() {
	fmt.Println("Application booting...")

	r := gin.Default()

	conn, err := database.GetConnection()

	if err != nil {
		panic(err)
	}

	//routing.Register(r, conn, os.Getenv("INSTANCE_ID"))
	routing.Register(r, conn, getAppContext())

	err = r.Run("0.0.0.0:8000")

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getAppContext() *common.ApplicationContext {
	//return common.NewApplicationContext(os.Getenv("INSTANCE_ID"), os.Getenv("APP_BASE_URL"))
	return common.NewApplicationContext("01", "http://localhost:8000/")
}

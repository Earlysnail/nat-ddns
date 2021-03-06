package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/account"
	"server/database"
	"server/frp"
	"server/middleware"
)

func main() {
	err := database.InitMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	err = database.InitRedis()
	if err != nil {
		log.Fatal(err)
	}

	//err = database.InitMysql()
	//if err != nil {
	//	log.Fatal(err)
	//}

	r := gin.Default()

	r.Use(
		middleware.CostTime(),
		middleware.Session(),
		middleware.SessionAuth(),
		middleware.Translation())

	account.Setup(r)
	frp.Setup(r)

	r.Run("0.0.0.0:9000")
}

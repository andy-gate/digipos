package main

import (
  "time"
  "fmt"
  "os"

  "github.com/gin-gonic/gin"

  "digipos/controllers"
  "digipos/models"
  "github.com/robfig/cron"

)

func main() {
  os.Setenv("TZ", "Asia/Jakarta")
  fmt.Printf("Started at : %3v \n", time.Now())

  //Init Redis
  // models.InitRedis()
  // defer models.RedisConn.Close()

  models.InitGormPostgres()
  defer models.MPosGORM.Close()

  // Set the router as the default one shipped with Gin
  gin.SetMode(gin.ReleaseMode)
  router := gin.Default()

  c := cron.New()
	// c.AddFunc("*/1 * * * *", func() { fmt.Println("[Job 1]Every minute job\n") })
	c.AddFunc("@hourly", func() { 
    fmt.Println("Cron started")
    controllers.GetHistoryPurchaseCronjob()
    controllers.GetHistoryDepositCronjob()
  })
  c.Start()

  // Setup route group for the API
  // api := router.Group("/api")

  // api.GET("/getHistoryPurchase", controllers.GetHistoryPurchase)
  // api.GET("/getHistoryDeposit", controllers.GetHistoryDeposit)
  
  // Start and run the server
  router.Run(":4000")
}
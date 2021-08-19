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

  models.InitGormPostgres()
  defer models.MPosGORM.Close()

  // Set the router as the default one shipped with Gin
  gin.SetMode(gin.ReleaseMode)
  router := gin.Default()

  c := cron.New()
	c.AddFunc("@hourly", func() { 
    fmt.Println("Cron started: %3v", time.Now())
    controllers.GetHistoryPurchaseCronjob()
    controllers.GetHistoryDepositCronjob()
  })
  c.Start()
  
  // Start and run the server
  router.Run(":4000")
}
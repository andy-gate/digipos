package main

import (
  "time"
  "fmt"
  "os"

  // "github.com/gin-gonic/gin"

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
  // gin.SetMode(gin.ReleaseMode)
  // router := gin.Default()
  // controllers.GetHistoryPurchaseCronjobES()

  c := cron.New()
	c.AddFunc("@hourly", func() { 
    fmt.Printf("Cron started at : %3v \n", time.Now())
    controllers.GetHistoryPurchaseCronjobES()
  })
  c.Start()
  
  for {
    time.Sleep(time.Second)
  }

  // Start and run the server
  // router.Run(":4000")
}
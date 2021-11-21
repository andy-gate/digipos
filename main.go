package main

import (
  "time"
  "fmt"
  "os"

  // "github.com/gin-gonic/gin"

  "digipos/controllers"
  "github.com/robfig/cron"

)

func main() {
  os.Setenv("TZ", "Asia/Jakarta")
  fmt.Printf("Started at : %3v \n", time.Now())

  // Set the router as the default one shipped with Gin
  // gin.SetMode(gin.ReleaseMode)
  // router := gin.Default()
  // controllers.GetHistoryPurchaseCronjobES()

  c := cron.New()
	c.AddFunc("@hourly", func() {
    controllers.GetHistoryPurchaseCronjobES()
  })
  c.AddFunc("0 20 * * * *", func() {
    controllers.GetHistoryPurchaseCronjobES2()
  })
  c.AddFunc("0 40 * * * *", func() {
    controllers.GetHistoryPurchaseCronjobES3()
  })
  c.Start()
  
  for {
    time.Sleep(time.Second)
  }

  // Start and run the server
  // router.Run(":4000")
}
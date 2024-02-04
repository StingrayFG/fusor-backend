package main

import (
  "fmt"
  "log"
  "os"

  //"github.com/teris-io/shortid"
  "github.com/joho/godotenv"

  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"

  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)


type Record struct {
  Uuid          string  `json:"uuid"`
  OriginalLink  string  `json:"originalLink"`
}


func getRecord(c *gin.Context) {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  db, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
  if err != nil {
    log.Fatal("Failed to connect database")
  }


  record := Record{}
  db.First(&record, "uuid = ?", c.Param("uuid"))
  
  fmt.Print(record)
  if (Record{} != record) {
    c.IndentedJSON(http.StatusOK, record)
  } else {
    c.IndentedJSON(http.StatusNotFound, Record{})
  }
}

func main() {
  router := gin.Default()

  router.ForwardedByClientIP = true
  router.SetTrustedProxies([]string{"127.0.0.1"})

  config := cors.DefaultConfig()
  config.AllowOrigins = []string{"http://localhost:4200"}
  router.Use(cors.New(config))
  

  router.GET("/record/get/:uuid", getRecord)

  router.Run("localhost:5200")
}
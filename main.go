package main

import (
  //"fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)

type Record struct {
  Uuid          string  `json:"uuid"`
  OriginalLink  string  `json:"originalLink"`
}

var rec = Record{Uuid: "1", OriginalLink: "1"}

func getRecord(c *gin.Context) {
  //fmt.Println(rec)
  c.IndentedJSON(http.StatusOK, rec)
}

func main() {
  router := gin.Default()
  router.ForwardedByClientIP = true
  router.SetTrustedProxies([]string{"127.0.0.1"})

  router.GET("/record/get", getRecord)

  router.Run("localhost:5200")
}
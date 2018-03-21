package main

import (
    "fmt"
    "os"

    "github.com/gin-gonic/gin"
)

type Person struct {
    ID        uint   `json:"id"`
    FirstName string `json:"firstname"`
    LastName  string `json:"lastname"`
}

func main() {
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()
    r.GET("/", handleRoot)
    //err := r.RunTLS("0.0.0.0:4443", "cert.pem", "key.pem") # enable TLS
    err := r.Run("0.0.0.0:8080")
    if err != nil {
        fmt.Fprintf(os.Stderr, err.Error())
        os.Exit(1)
    }
}

func handleRoot(c *gin.Context) {
    person1 := &Person{ID: 111, FirstName: "Daffy", LastName: "Duck"}
    person2 := &Person{ID: 112, FirstName: "Goofy", LastName: "Dog"}
    x := []*Person{person1, person2}
    c.IndentedJSON(200, x)
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tmilewski/goenv"
)

var isProduction bool

func main() {
	checkErr(goenv.Load())

	base := GetBase()
	r := gin.Default()

	r.GET("/users", base.listUsers)
	r.GET("/users/:user_id", base.getUser)
	r.DELETE("/users/:user_id", base.deleteUser)
	r.PUT("/users/:user_id", base.editUser)
	r.POST("/users", base.addUser)

	r.Run(":8080")
}

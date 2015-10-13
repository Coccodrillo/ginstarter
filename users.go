package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosexy/to"
	"gopkg.in/gorp.v1"
)

func (b *Base) listUsers(c *gin.Context) {
	var users []User
	_, err := b.Db.Select(&users, "SELECT * FROM users ORDER BY id DESC")
	checkErr(err)
	if err == nil {
		c.JSON(200, users)
		return
	}
	c.JSON(404, "Not found")
}

func (b *Base) getUser(c *gin.Context) {
	if userId := to.Int64(c.Params.ByName("user_id")); userId != 0 {
		var user User
		err := b.Db.SelectOne(&user, "SELECT * FROM users WHERE id = ? LIMIT 1", userId)
		checkErr(err)
		if err == nil {
			c.JSON(200, user)
			return
		}
	}
	c.JSON(404, "Not found")
}

func (b *Base) editUser(c *gin.Context) {
	var jsonForm UserJSON
	if err := c.BindJSON(&jsonForm); err != nil {
		c.JSON(400, "Bad request")
		return
	}
	if userId := to.Int64(c.Params.ByName("user_id")); userId != 0 {
		user := User{Id: userId, Name: jsonForm.Name, Email: jsonForm.Email}
		_, err := b.Db.Update(&user)
		checkErr(err)
		c.JSON(200, user)
		return
	}
	c.JSON(404, "Not found")
}

func (b *Base) addUser(c *gin.Context) {
	var jsonForm UserJSON
	if err := c.BindJSON(&jsonForm); err != nil {
		c.JSON(400, "Bad request")
		return
	}
	user := User{Name: jsonForm.Name, Email: jsonForm.Email}
	err := b.Db.Insert(&user)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(201, user)
}

func (b *Base) deleteUser(c *gin.Context) {
	if userId := to.Int64(c.Params.ByName("user_id")); userId != 0 {
		var user User
		err := b.Db.SelectOne(&user, "SELECT * FROM users WHERE id = ? LIMIT 1", userId)
		checkErr(err)
		if err == nil {
			_, err = b.Db.Delete(&user)
			checkErr(err)
			c.Writer.WriteHeader(204)
			return
		}
	}
	c.JSON(400, "Bad request")
}

func (u *User) PreInsert(s gorp.SqlExecutor) error {
	u.RegistrationDate = time.Now().Unix()
	return nil
}

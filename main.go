package main

import (
	"k8s-mig-srv-go/user"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=k8s-mig-db-old user=test001 password=test001 dbname=msghub port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	router := gin.Default()
	v0 := router.Group("/srv/go")
	{
		v0.GET("/ping", func(c *gin.Context) {
			c.SecureJSON(http.StatusOK, "pong")
			return
		})
		v0.GET("/users", func(c *gin.Context) {
			query := user.User{}

			err := c.ShouldBindQuery(&query)
			if err != nil {
				c.SecureJSON(http.StatusBadRequest, err.Error())
				return
			}

			if (user.User{} == query) {
				c.SecureJSON(http.StatusOK, "please input search criteria")
				return
			}

			result, err := user.GetUsers(db, query)
			if err != nil {
				c.SecureJSON(http.StatusBadRequest, err.Error())
				return
			}

			c.SecureJSON(http.StatusOK, result)
			return
		})
		v0.GET("/users/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.SecureJSON(http.StatusBadRequest, err.Error())
				return
			}

			result, err := user.GetUser(db, id)
			if err != nil {
				c.SecureJSON(http.StatusBadRequest, err.Error())
				return
			}

			c.SecureJSON(http.StatusOK, result)
			return
		})
		v0.POST("/users", func(c *gin.Context) {
			query := user.User{}

			err := c.ShouldBindJSON(&query)
			if err != nil {
				c.SecureJSON(http.StatusBadRequest, err.Error())
				return
			}

			if (user.User{} == query) {
				c.SecureJSON(http.StatusOK, "please input search criteria")
				return
			}

			success, err := user.CreateUser(db, query)
			if err != nil {
				c.SecureJSON(http.StatusBadRequest, err.Error())
				return
			}

			c.SecureJSON(http.StatusOK, success)
			return
		})
		v0.PUT("/users/:id", func(c *gin.Context) {
			query := user.User{}

			err := c.ShouldBindJSON(&query)
			if err != nil {
				c.SecureJSON(http.StatusBadRequest, err.Error())
				return
			}

			if (user.User{} == query) {
				c.SecureJSON(http.StatusOK, "please input search criteria")
				return
			}

			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.SecureJSON(http.StatusBadRequest, err.Error())
				return
			}

			success, err := user.UpdateUser(db, query, id)
			if err != nil {
				c.SecureJSON(http.StatusBadRequest, err.Error())
				return
			}

			c.SecureJSON(http.StatusCreated, success)
			return
		})
	}

	router.Run(":3000")
}

package main

import (
	"k8s-mig-srv-go/user"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	viper.AutomaticEnv()

	prepareDBConnection()

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

func prepareDBConnection() {
	var builder strings.Builder

	builder.WriteString("host=")
	builder.WriteString(viper.GetString("DB_OLD_HOST"))
	builder.WriteString(" user=")
	builder.WriteString(viper.GetString("DB_OLD_USER"))
	builder.WriteString(" password=")
	builder.WriteString(viper.GetString("DB_OLD_PASSWORD"))
	builder.WriteString(" dbname=")
	builder.WriteString(viper.GetString("DB_OLD_NAME"))
	builder.WriteString(" port=")
	builder.WriteString(viper.GetString("DB_OLD_PORT"))
	builder.WriteString(" sslmode=disable")

	db, _ = gorm.Open(postgres.Open(builder.String()), &gorm.Config{})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

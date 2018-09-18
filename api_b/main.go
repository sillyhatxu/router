package main

import (
	log "github.com/xushikuan/microlog"
	consulClient "github.com/xushikuan/consul-client"
	"github.com/gin-gonic/gin"
	"strconv"
	"golang-cloud/tool/http/http_response"
	"net/http"
	"flag"
)

func main() {
	router := gin.Default()

	enviroment := flag.String("enviroment", "", "a string")
	consulAddress := flag.String("consul-address", "", "a string")
	moduleName := flag.String("module-name", "", "a string")
	port := flag.Int("port", 80, "an int")
	flag.Parse()

	log.Info("consulAddress : ",*consulAddress)
	log.Info("moduleName : ",*moduleName)
	log.Info("port : ",*port)
	group := router.Group("/user")
	{
		group.POST("/created", testCreatedAPI)
		group.PUT("/updated/:id", testUpdatedAPI)
		group.GET("/get/:id", testGetAPI)
		group.DELETE("/deleted/:id", testDeletedAPI)
	}
	healthAPI :="/health"
	router.GET(healthAPI, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Description":"Golang project client status",
			"Status":"UP",
		})
	})
	consulServer := consulClient.NewConsulServer(*enviroment,*moduleName,*port,healthAPI)
	consulServer.SetConsulAddress(*consulAddress)
	consulServer.RegisterServer()
	router.Run(":" + strconv.Itoa(*port))
}

func testCreatedAPI(context *gin.Context)  {
	var requestBody *User
	if err := context.ShouldBindJSON(&requestBody); err == nil {
		context.JSON(http.StatusOK, http_response.Success("BBBBB----- Created : {Name:" + requestBody.Name+ ",Age:" + strconv.Itoa(requestBody.Age) + "}"))
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func testUpdatedAPI(context *gin.Context)  {
	id := context.Param("id")
	var requestBody *User
	if err := context.ShouldBindJSON(&requestBody); err == nil {
		context.JSON(http.StatusOK, http_response.Success("BBBBB----- Updated : {id:" + id + ",Name:" + requestBody.Name+ ",Age:" + strconv.Itoa(requestBody.Age) + "}"))
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func testGetAPI(context *gin.Context)  {
	id := context.Param("id")
	context.JSON(http.StatusOK, http_response.Success("BBBBB----- Get : "+id))
}

func testDeletedAPI(context *gin.Context)  {
	id := context.Param("id")
	context.JSON(http.StatusOK, http_response.Success("BBBBB----- Deleted : "+id))
}

type User struct {
	Name string `form:"name" binding:"required"`
	Age  int    `form:"age" binding:"required"`
}
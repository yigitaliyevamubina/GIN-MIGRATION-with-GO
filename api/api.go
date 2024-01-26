package main

import (
	"GIN_MIGRATION/models"
	"GIN_MIGRATION/storage"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/user/create", CreateUser)

	router.PUT("/user/update/:id", UpdateUser)

	router.DELETE("/user/delete/:id", DeleteUser)

	router.GET("/user/get/:id", GetUserById)

	router.GET("user/get/all", GetAllUsers)

	router.GET("user/filter/name", FilterByName)

	router.GET("/user/get/role", GetUsersByRoleHandler)

	if err := router.Run("localhost:8090"); err != nil {
		log.Println(err)
	}
}

func CreateUser(c *gin.Context) {
	var reqUser models.User

	if err := c.BindJSON(&reqUser); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	respUser, err := storage.CreateUser(reqUser)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusCreated, respUser)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	var reqUser models.User
	if err := c.BindJSON(&reqUser); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	updatedUser, err := storage.UpdateUser(intId, reqUser)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusFound, updatedUser)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	deletedUser, err := storage.DeleteUser(intId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusFound, deletedUser)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	respUser, err := storage.GetUserById(intId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusFound, respUser)
}

func GetAllUsers(c *gin.Context) {
	limit := c.Query("limit")

	convertedLimit, err := strconv.Atoi(limit)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	page := c.Query("page")

	convertedPage, err := strconv.Atoi(page)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	respUsers, err := storage.GetAllUsers(convertedLimit, convertedPage)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusOK, respUsers)
}

func FilterByName(c *gin.Context) {
	name := c.Query("name")
	limit := c.Query("limit")

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	page := c.Query("page")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	respUsers, err := storage.FilterByName(name, intLimit, intPage)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusOK, respUsers)
}

func GetUsersByRoleHandler(c *gin.Context) {
	roleId, err := strconv.Atoi(c.Query("role"))
	fmt.Println(roleId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	fmt.Println(limit)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	page, err := strconv.Atoi(c.Query("page"))
	fmt.Println(page)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respUser, err := storage.GetUsersByRole(roleId, limit, page)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, respUser)
}

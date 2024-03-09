package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_api_deploy/database"
	"rest_api_deploy/repository"
	"rest_api_deploy/structs"
	"strconv"
)

var(
	DB = database.DBSetup()
)

func GetAllPerson(c *gin.Context) {
	var(
		result gin.H
	)

	persons, err := repository.GetAllPerson(DB)

	if err!= nil {
		result =gin.H {
			"result": err,
		}
	} else {
		result =gin.H {
			"result": persons,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertPerson(c *gin.Context) {
	var person structs.Person

	err := c.ShouldBindJSON(&person)
	if err != nil {
		panic(err)
	}

	err = repository.InsertPerson(DB, person)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H {
		"result": "Insert Person Success",
	})
}

func UpdatePerson(c *gin.Context) {
	var person structs.Person
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&person)
	if err != nil {
		panic(err)
	}

	person.ID = int64(id)

	err = repository.InsertPerson(DB, person)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H {
		"result": "Update Person Success",
	})
}

func DeletePerson(c *gin.Context) {
	var person structs.Person
	id, err := strconv.Atoi(c.Param("id"))

	person.ID = int64(id)

	err = repository.InsertPerson(DB, person)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H {
		"result": "Delete Person Success",
	})
}
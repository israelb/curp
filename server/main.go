package main

import (
	"net/http"

	"github.com/israelb/curp"

	"github.com/gin-gonic/gin"
)

var msg struct {
	Names          string `json:"names"`
	FirstLastName  string `json:"firstLastName"`
	SecondLastName string `json:"secondLastName"`
	Sex            string `json:"sex"`
	StateCode      string `json:"stateCode"`
	BirthDate      string `json:"birthDate"`
}

func main() {
	router := gin.Default()

	router.GET("/curp", createCurp)

	router.Run(":1700")
}

func createCurp(c *gin.Context) {
	msg.Names = c.Query("names")
	msg.FirstLastName = c.Query("firstLastName")
	msg.SecondLastName = c.Query("secondLastName")
	msg.Sex = c.Query("sex")
	msg.StateCode = c.Query("stateCode")
	msg.BirthDate = c.Query("birthDate")

	message, errMessage := curp.NewCurp(msg.Names, msg.FirstLastName, msg.SecondLastName, msg.Sex, msg.StateCode, msg.BirthDate)

	if errMessage != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMessage.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"person": msg, "curp": message, "status": http.StatusOK})
	}
}

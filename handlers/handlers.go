package handlers

import (
	"DemoHttpMock2/models"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ApiUrl = "https://reqres.in/api/users?page=2"

func CreateUser(c *gin.Context) ([]byte, error) {

	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return []byte{}, err
	}

	//c.JSON(http.StatusOK, gin.H{"user": user})
	return []byte("success"), nil
}

func GetCustomData(c *gin.Context) ([]byte, error) {

	req, err := http.NewRequest(http.MethodGet, ApiUrl, nil)
	if err != nil {
		fmt.Println("Error buiding request object")
		return []byte{}, err
	}

	client := &http.Client{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request")
		return []byte{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error response from Api")
		return []byte{}, err
	}

	return body, nil
}

func RequestWrapper(f func(c *gin.Context) ([]byte, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := f(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	}
}

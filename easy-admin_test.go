package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nicelizhi/easy-admin/app/admin/apis"
	"github.com/stretchr/testify/assert"
)

func HomepageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome golang"})
}

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	fmt.Println(router.Routes())
	return router
}
func TestHomepageHandler(t *testing.T) {
	// mockResponse := `{"message":"Welcome goland"}`
	r := SetUpRouter()
	r.GET("/", apis.EasyAdminStart)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// responseData, _ := ioutil.ReadAll(w.Body)
	// assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserLoginHandler(t *testing.T) {
	// r := SetUpRouter()
	// r.GET("/api/v1/getinfo", apis.SysUser.GetInfo)

}

package openapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)
func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
// Mock data for testing
var mockMenuItems = []MenuItem{
	{ID: 1, Name: "Pizza", Price: 10.99},
	{ID: 2, Name: "Burger", Price: 7.49},
}

// Mock handler for ShowMenuItemById
func GetMockMenuItemById(c *gin.Context) {
	idParam := c.Param("itemId")
	id, _ := strconv.Atoi(idParam)

	for _, item := range mockMenuItems {
		if int(item.ID) == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "menu item not found"})
}

func TestShowMenuItemByIdHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/testmenuitem/:itemId", GetMockMenuItemById)

	// Test case: Get a valid menu item by ID
	req, _ := http.NewRequest("GET", "/testmenuitem/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var menuItem MenuItem
	json.Unmarshal(w.Body.Bytes(), &menuItem)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Pizza", menuItem.Name)
	assert.Equal(t, 10.99, menuItem.Price)

	// Test case: Get an invalid menu item (non-existent ID)
	req, _ = http.NewRequest("GET", "/testmenuitem/999", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
// package openapi

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func SetUpRouter() *gin.Engine {
// 	router := gin.Default()
// 	return router
// }

// func TestGetMenuItemsHandler(t *testing.T) {
// 	r := SetUpRouter()
// 	r.GET("/testmenuitems", GetTestMenuItems)
// 	req, _ := http.NewRequest("GET", "/testmenuitems", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	var menuitems []MenuItem
// 	json.Unmarshal(w.Body.Bytes(), &menuitems)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.NotEmpty(t, menuitems)
// }

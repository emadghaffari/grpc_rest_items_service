package item

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/emadghaffari/grpc_rest_items_service/model/items"
	"github.com/emadghaffari/grpc_rest_items_service/model/queries"
	"github.com/emadghaffari/grpc_rest_items_service/services"
	"github.com/emadghaffari/res_errors/errors"
	"github.com/gin-gonic/gin"
)

// Create func
func Create(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.Request.Header.Get("user_id"),10,64)
	
	responseBody,err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		resErr := errors.HandlerBadRequest("invalid Body request")
		c.JSON(resErr.Status(), gin.H{"message":resErr.Message(),"causes": resErr.Causes()})
		return
	}
	defer c.Request.Body.Close()

	var item items.Item
	if err := json.Unmarshal(responseBody,&item); err !=nil {
		resErr := errors.HandlerBadRequest("invalid Body request we can not unmarshal the request")
		c.JSON(resErr.Status(), gin.H{"message":resErr.Message(),"causes": resErr.Causes()})
		return
	}
	item.Seller = userID

	result,resErr := services.ItemService.Create(item)
	if resErr != nil {
		c.JSON(resErr.Status(), gin.H{"message":resErr.Message(),"causes": resErr.Causes()})
		return
	}
	c.JSON(http.StatusOK, result)
}
// Get func
func Get(c *gin.Context) {
	itemID := c.Param("id")
	result,resErr := services.ItemService.Get(itemID)
	if resErr != nil {
		c.JSON(resErr.Status(), gin.H{"message":resErr.Message(),"causes": resErr.Causes()})
		return
	}
	c.JSON(http.StatusOK, result)

}
// Search func
func Search(c *gin.Context) {
	responseBody,err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		resErr := errors.HandlerBadRequest("invalid Body request")
		c.JSON(resErr.Status(), gin.H{"message":resErr.Message(),"causes": resErr.Causes()})
		return
	}
	defer c.Request.Body.Close()

	var query queries.EsQuery
	if err := json.Unmarshal(responseBody, &query); err !=nil {
		resErr := errors.HandlerBadRequest("invalid Body request we can not unmarshal the request")
		c.JSON(resErr.Status(), gin.H{"message":resErr.Message(),"causes": resErr.Causes()})
		return
	}

	result,resErr := services.ItemService.Search(query)
	if resErr != nil {
		c.JSON(resErr.Status(), gin.H{"message":resErr.Message(),"causes": resErr.Causes()})
		return
	}

	c.JSON(http.StatusOK, result)

}
// Delete func
func Delete(c *gin.Context) {
	itemID := c.Param("id")
	resErr := services.ItemService.Delete(itemID)
	if resErr != nil {
		c.JSON(resErr.Status(), gin.H{"message":resErr.Message(),"causes": resErr.Causes()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status":"success"})

}

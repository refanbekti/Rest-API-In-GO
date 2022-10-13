package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Items struct {
	ItemId      string `json:"item_id"`
	ItemCode    string `json:"itemcode"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
}

var ItemsDatas = []Items{}

func CreateOrders(ctx *gin.Context) {

	var newItems Items

	if err := ctx.ShouldBindJSON(&newItems); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return

	}

	newItems.ItemId = fmt.Sprintf("%d", len(ItemsDatas)+1)

	ItemsDatas = append(ItemsDatas, newItems)

	ctx.JSON(http.StatusCreated, gin.H{

		"Items": newItems,
	})
}

func GetItems(ctx *gin.Context) {

	ItemId := ctx.Param("ordersID")
	condition := false
	var itemsData Items

	for i, items := range ItemsDatas {

		if ItemId == items.ItemId {

			condition = true

			itemsData = ItemsDatas[i]

			break

		}
	}

	if !condition {

		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", ItemId),
		})

		return

	}

	ctx.JSON(http.StatusOK, gin.H{
		"Items": itemsData,
	})

}

func UpdateItems(ctx *gin.Context) {
	itemsID := ctx.Param("orderid")
	condition := false
	var updatedItems Items

	if err := ctx.ShouldBindJSON(&updatedItems); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, items := range ItemsDatas {
		if itemsID == items.ItemId {
			condition = true
			ItemsDatas[i] = updatedItems

			break

		}
	}

	if !condition {

		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", itemsID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id Sv has been successfully updated", itemsID),
	})
}

func DeleteItems(ctx *gin.Context) {

	itemsID := ctx.Param("orderid")
	condition := false
	var itemsIndex int

	for i, items := range ItemsDatas {

		if itemsID == items.ItemId {
			condition = true
			itemsIndex = i
			break
		}

	}

	if !condition {

		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id v not found", itemsID),
		})

		return
	}

	copy(ItemsDatas[itemsIndex:], ItemsDatas[itemsIndex+1:])
	ItemsDatas[len(ItemsDatas)-1] = Items{}
	ItemsDatas = ItemsDatas[:len(ItemsDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully deleted", itemsID),
	})

}

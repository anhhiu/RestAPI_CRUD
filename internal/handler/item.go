package handler

import (
	"net/http"
	"strconv"

	"RestApi_CRUD/internal/connect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HoTen struct {
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Address string `json:"address"`
}

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var ht HoTen
		if err := c.ShouldBindJSON(&ht); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		per := connect.Persons{
			Name:         ht.Name,
			Gender:       ht.Gender,
			Address:      ht.Address,
			Phone_number: "unknown",
		}

		err := connect.AddUser(db, per)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": per,
		})
	}
}

func GetAllItems(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		users, err := connect.GetAllUsers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": users,
		})
	}
}

func GetItemByID(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID format",
			})
			return
		}

		user, err := connect.FindUserByID(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}
}

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var ht HoTen
		if err := c.ShouldBindJSON(&ht); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID format",
			})
			return
		}

		user, err := connect.FindUserByID(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		user.Name = ht.Name
		user.Gender = ht.Gender
		user.Address = ht.Address

		err = connect.UpdateUser(db, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}
}

func DeleteItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID format",
			})
			return
		}

		err = connect.DeleteUser(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Xóa thành công",
		})
	}
}

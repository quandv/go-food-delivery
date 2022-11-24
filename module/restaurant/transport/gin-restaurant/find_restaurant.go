package ginrestaurant

import (
	restaurantbusiness "food-delivery/module/restaurant/business"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindRestaurantById(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "id must be an integer",
			})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		business := restaurantbusiness.NewFindRestaurantBusiness(store)

		result, err := business.FindOne(ctx, map[string]interface{}{"id": id})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": &result,
		})
	}
}

func FindRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, err := strconv.Atoi(ctx.Param("page"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "page must be an integer",
			})
			return
		}

		limit, err := strconv.Atoi(ctx.Param("limit"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "limit must be an integer",
			})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		business := restaurantbusiness.NewFindRestaurantBusiness(store)

		result, err := business.Find(ctx, map[string]interface{}, map[string]interface{}{
			"page":  page,
			"limit": limit,
		})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": &result,
		})
	}
}

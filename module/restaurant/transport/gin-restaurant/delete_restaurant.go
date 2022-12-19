package ginrestaurant

import (
	common "go-food-delivery/common"
	restaurantbusiness "go-food-delivery/module/restaurant/business"
	restaurantstorage "go-food-delivery/module/restaurant/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteRestaurantById(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.InvalidRequest(err, "id must be an integer"))
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		business := restaurantbusiness.NewDeleteRestaurantBusiness(store)

		if err := business.DeleteOne(ctx, map[string]interface{}{"id": id}); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Delete restaurant successful",
		})
	}
}

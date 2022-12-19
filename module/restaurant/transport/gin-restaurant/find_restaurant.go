package ginrestaurant

import (
	common "go-food-delivery/common"
	restaurantbusiness "go-food-delivery/module/restaurant/business"
	restaurantmodel "go-food-delivery/module/restaurant/model"
	restaurantstorage "go-food-delivery/module/restaurant/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindRestaurantById(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.InvalidRequest(err, "id must be an integer"))
			return
		}

		// Test middleware.Recover() catch panicking
		// var arr []string
		// log.Println(arr[0])

		// Panicking in other goroutine
		go func() {
			// Nếu không sd recover ở đây => crash app
			// Bởi vì goroutine này độc lập với goroutine đang được sử dụng trong main.go
			defer func() {
				if r := recover(); r != nil {
					log.Println("Recover inside other goroutine: ", r)
				}
			}()

			var arr []string
			log.Println(arr[0])
		}()

		store := restaurantstorage.NewSQLStore(db)
		business := restaurantbusiness.NewFindRestaurantBusiness(store)

		result, err := business.FindOne(ctx, map[string]interface{}{"id": id})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err) // tại sao lại truyền trực tiếp "err" mà ko wrap bằng common error
			// => bởi vì error này đã được wrap bằng common error trong storage layer
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": &result,
		})
	}
}

func FindRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Pagination

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, common.InvalidRequest(err, ""))
			return
		}

		paging.FullFill()

		var filter restaurantmodel.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, common.InvalidRequest(err, ""))
		}

		store := restaurantstorage.NewSQLStore(db)
		business := restaurantbusiness.NewFindRestaurantBusiness(store)

		result, err := business.Find(ctx, filter, &paging)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.NewResponse(result, paging, filter))
	}
}

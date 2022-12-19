package ginrestaurant

import (
	appctx "go-food-delivery/component/app-context"
	restaurantbusiness "go-food-delivery/module/restaurant/business"
	restaurantmodel "go-food-delivery/module/restaurant/model"
	restaurantstorage "go-food-delivery/module/restaurant/storage"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		db := appCtx.GetMainDB()

		log.Println("data: ", data)

		if err := ctx.ShouldBind(&data); err != nil {
			log.Println("ctx.ShouldBind => ", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		business := restaurantbusiness.NewCreateRestaurantBusiness(store)

		log.Println(ctx.Request.Context())

		if err := business.CreateRestaurant(ctx.Request.Context(), &data); err != nil {
			log.Println("business.CreateRestaurant => ", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Created new restaurant successful",
			"data":    &data,
		})
	}
}

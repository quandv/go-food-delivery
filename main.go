package main

import (
	appctx "go-food-delivery/component/app-context"
	middleware "go-food-delivery/middleware"
	"go-food-delivery/module/restaurant/transport/gin-restaurant"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("FOOD_DELIVERY_DB_CONNECT_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	appCtx := appctx.NewAppContext(db)

	// === CREATE NEW RESTAURANT ===
	// newRestaurantRecord := Restaurant{
	// 	Name:    "Ha noi beer",
	// 	Address: "219 Trung Kinh, Cau Giay, Ha Noi",
	// }
	// insertNewRestaurant := db.Create(&newRestaurantRecord)

	// if insertNewRestaurant.Error != nil {
	// 	log.Fatalln(insertNewRestaurant.Error) // trong log module có các method: Print(f|ln) || Panic(f|ln) || Fatal(f|ln)
	// 	// log.Fatal(f|ln) sau khi in message ra ngoài màn hình thì nó sẽ exit program với code = 1 (bằng việc call os.Exit(1))
	// }

	// log.Println(insertNewRestaurant.RowsAffected)

	// === UPDATE A RESTAURANT ===
	// Best practice là tạo ra một struct mới để sử dụng cho update, ko dùng chung struct Restaurant
	// Bởi vì nếu use Restaurant thì một số giá trị truyền vào để update sẽ bị gorm duyệt và bỏ qua không cho update: zeroed values (0 | "" | false)
	// => các types sẽ phải là dạng con trỏ, bởi vì dạng con trỏ sẽ chỉ được gorm duyệt qua và loại bỏ các giá trị nil, ko loại bỏ zeroed values

	// updateRestaurantName := "Sai gon beer 2"
	// updateRestaurantRecord := RestaurantUpdate{
	// 	Name: &updateRestaurantName,
	// }

	// // updateRestaurant := db.Model(Restaurant{}).Where("id = ?", 4).Updates(updateRestaurantRecord)
	// updateRestaurant := db.Model(&updateRestaurantRecord).Where("id = ?", 4).Updates(updateRestaurantRecord)

	// if updateRestaurant.Error != nil {
	// 	log.Fatalln(updateRestaurant.Error)
	// }

	// log.Println(updateRestaurant.RowsAffected)

	// === GET A RESTAURANT ===
	// getFirstRestaurantWithCondition := db.First(&Restaurant{}, "id = ?", "5")
	// log.Println(getFirstRestaurantWithCondition)

	// firstRestaurant := db.First(&Restaurant{})
	// log.Println(firstRestaurant)

	// lastRestaurant := db.Last(&Restaurant{})
	// log.Println(lastRestaurant)

	// randRestaurant := db.Take(&Restaurant{})
	// log.Println(randRestaurant)

	// lastRestaurant := Restaurant{}
	// db.Last(&lastRestaurant)

	// log.Println(lastRestaurant)

	// === DELETE A RESTAURANT ===
	// db.Where("id = ?", 5).Delete(&Restaurant{})

	// getDeletedRestaurant := db.First(&Restaurant{}, "id = ?", "5")

	// type User struct {
	// 	gorm.Model
	// 	Name         string
	// 	CompanyRefer int
	// 	Company      Company `gorm:"foreignKey:CompanyRefer"`
	// 	// use CompanyRefer as foreign key
	// }

	// Create web server by Gin framework
	app := gin.Default()
	app.Use(middleware.Recover(appCtx))
	restaurant := app.Group("/restaurants")

	restaurant.POST("", ginrestaurant.CreateRestaurant(appCtx))

	// chưa chuyển hết các endpoints dưới để sử dụng appCtx => mỗi ngày vào sửa 1 endpoint cho nhớ phần app context
	restaurant.GET("/:id", ginrestaurant.FindRestaurantById(db))
	restaurant.GET("/", ginrestaurant.FindRestaurant(db))
	restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurantById(db))

	app.Run(":8888") // listen and serve on 0.0.0.0:8888
}

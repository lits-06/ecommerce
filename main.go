package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lits-06/ecommerce/controllers"
	"github.com/lits-06/ecommerce/database"
	"github.com/lits-06/ecommerce/middleware"
	"github.com/lits-06/ecommerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Product"), 
									  database.UserData(database.Client, "User"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}

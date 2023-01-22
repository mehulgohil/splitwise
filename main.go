package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/splitwise/controller"
	"github.com/mehulgohil/splitwise/service"
	"log"
)

func main()  {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "splitwise",
	}
	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	serviceStruct := service.ServiceStruct{
		DB: db,
	}
	handlerStruct := controller.HandlerStruct{
		ServiceStruct: serviceStruct,
	}

	app := iris.New()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
	})

	app.UseRouter(crs)

	app.Post("/transactions", handlerStruct.PostTransactionHandler)
	app.Get("/transactions/oweBy/{mobileNo: string}", handlerStruct.GetOweByTransactionHandler)
	app.Get("/transactions/oweTo/{mobileNo: string}", handlerStruct.GetOweToTransactionHandler)
	app.Patch("/transactions/{transactionId: int}", handlerStruct.PatchTransactionHandler)

	_ = app.Listen(":8080")

}

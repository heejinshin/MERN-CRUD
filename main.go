package main

import (
	"fmt"

	"github.com/jackbalageru/MERN-CRUD/go-server/pkg/api"
	"github.com/jackbalageru/MERN-CRUD/go-server/pkg/db"
	"github.com/jackbalageru/MERN-CRUD/go-server/pkg/router"
)

func main() {
	
	dbHandler, err := db.NewAndConnectGorm("user:password@tcp(127.0.0.1:3307)/dev?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	apis := api.NewAPI(dbHandler)  // DB공간에 POST요청 db를 저장시킴 
	r := router.Router(apis)

	r.Run(":8081")
}



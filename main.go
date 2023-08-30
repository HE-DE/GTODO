package main

import (
	"GTODO/router"
	"GTODO/utils"
)

func main() {
	utils.InitConfig()
	utils.InitDB()
	r := router.Router()
	err := r.Run(":8888")
	if err != nil {
		panic(err)
	}
}

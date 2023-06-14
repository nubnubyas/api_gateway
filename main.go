package main

import (
	api "github.com/cloudwego/api_gateway/kitex_gen/api/studentapi"
	"log"
)

func main() {
	svr := api.NewServer(new(StudentApiImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

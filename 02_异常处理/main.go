package main

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/liyanfeng123/golang-promotion/02_异常处理/db"
	"github.com/liyanfeng123/golang-promotion/02_异常处理/service"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", GetUserHandler)
	if err := http.ListenAndServe(":8090", mux); err != nil {
		log.Fatal(err)
	}
}

func GetUserHandler(writer http.ResponseWriter, request *http.Request){
	configService := &service.IssueConfigService{}
	data, err := configService.GetTestData(context.Background())
	if err != nil {
		fmt.Fprintf(writer,err.Error())
		return
	}
	res,err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(writer,err.Error())
		return
	}
	fmt.Fprintf(writer,string(res))
}
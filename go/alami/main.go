package main

import (
	"alami/service"
	"log"
)

func main() {

	svc := service.NewServiceCore()
	listData := svc.GetBankData("Before Eod.csv")

	svc.CreateCSV(listData)

	log.Println("success generate after eod file !!!")

}

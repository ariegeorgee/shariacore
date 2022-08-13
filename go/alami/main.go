package main

import (
	"alami/service"
)

func main() {

	svc := service.NewServiceCore()
	listData := svc.GetBankData("Before Eod.csv")

	svc.CreateCSV(listData)

}

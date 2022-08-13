package service

import (
	"alami/model"
	"alami/repository"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type CoreService struct {
}

func NewServiceCore() repository.Process {
	return &CoreService{}
}

func (CoreService) GetBankData(path string) []*model.Nasabah {
	filenya, err := os.Open(path)

	if err != nil {
		fmt.Println("Cannot read file csv")
	}

	csvData, err := csv.NewReader(filenya).ReadAll()
	if err != nil {
		fmt.Println("something error")
		return []*model.Nasabah{}
	}

	var nasabahList []*model.Nasabah
	var waitGroup1 sync.WaitGroup
	var waitGroup2 sync.WaitGroup

	// thread2 := math.Floor(100 / 8)
	thread2 := 8

	waitGroup2.Add(thread2)
	thNum := 1

	fmt.Println("Thred 2 : ", thread2)

	for i, rowData := range csvData {
		if i == 0 {
			continue
		}

		nasabah := new(model.Nasabah)
		tmp := strings.Split(rowData[0], ";")
		nasabah.Setter(tmp)

		waitGroup1.Add(3)
		go func() { // No 1 Thread-No
			nasabah.SumAverageBalanced()
			nasabah.SetNo1()
			waitGroup1.Done()
		}()

		go func() {
			nasabah.AddBonusBalance() // No 2b Thread-No
			nasabah.SetNo2b()
			waitGroup1.Done()
		}()

		go func() {
			nasabah.FreeTransferCount() // No 2a Thread-No
			nasabah.SetNo2a()
			waitGroup1.Done()
		}()

		go func(num int) {
			if num < 100 {
				nasabah.AddLuckyBonusBalance() // No 3 Thread-No
				thD := num + 1
				nasabah.SetNo3(thNum)
				if thD%12 == 0 {
					fmt.Println("MODULUS DONE : ", thNum)
					thNum += 1
					waitGroup2.Done()
				}
			}

		}(i)

		nasabahList = append(nasabahList, nasabah)
	}

	defer filenya.Close()
	waitGroup1.Wait()
	waitGroup2.Wait()
	return nasabahList
}

func (CoreService) CreateCSV(nasabahList []*model.Nasabah) {
	csvFile, err := os.Create("Nasabah.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
		return
	}

	j, _ := json.Marshal(nasabahList)
	log.Println(string(j))

	w := csv.NewWriter(csvFile)
	w.Comma = ';'

	if err := w.Write([]string{
		"id",
		"Nama",
		"Age",
		"Balanced",
		"No 2b Thread-No",
		"No 3 Thread-No",
		"Previous Balanced",
		"Average Balanced",
		"No 1 Thread-No",
		"Free Transfer",
		"No 2a Thread-No",
	}); err != nil {
		log.Fatalln("error writing record to file", err)
		fmt.Println(err.Error())
		log.Println(err.Error())
		return
	}

	for _, v := range nasabahList {
		row := []string{
			fmt.Sprint(v.Id),
			fmt.Sprint(v.Nama),
			fmt.Sprint(v.Age),
			fmt.Sprint(v.Balanced),
			fmt.Sprint(v.No2b),
			fmt.Sprint(v.No3),
			fmt.Sprint(v.PreviousBalanced),
			fmt.Sprint(v.AverageBalanced),
			fmt.Sprint(v.No1),
			fmt.Sprint(v.FreeTransfer),
			fmt.Sprint(v.No2a),
		}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
			fmt.Println(err.Error())
			log.Println(err.Error())
			return
		}

	}

	w.Flush()
	defer csvFile.Close()
}

package model

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

type Nasabah struct {
	Id               int     `csv:"id"`
	Nama             string  `csv:"Nama"`
	Age              int     `csv:"Age"`
	Balanced         float64 `csv:"Balanced"`
	PreviousBalanced float64 `csv:"Previous Balanced"`
	AverageBalanced  float64 `csv:"Average Balanced"`
	FreeTransfer     float64 `csv:"Free Transfer"`
	No1              int
	No2a             int
	No2b             int
	No3              int
}

var NasabahList []*Nasabah
var err error

func (n *Nasabah) SumAverageBalanced() {
	n.AverageBalanced = (n.Balanced + n.PreviousBalanced) / 2
}

func (n *Nasabah) FreeTransferCount() {
	if n.Balanced >= 100 && n.Balanced <= 150 { // No 2a Thread-No
		n.FreeTransfer = 5
	}
}

func (n *Nasabah) AddBonusBalance() {
	if n.Balanced >= 150 { // No 2a Thread-No
		n.Balanced += 25
	}
}

func (n *Nasabah) AddLuckyBonusBalance() {
	n.Balanced += 10
}

func (n *Nasabah) SetId(tmp []string) {
	n.Id, err = strconv.Atoi(tmp[0])
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
}

func (n *Nasabah) SetName(tmp []string) {
	n.Nama = tmp[1]
}

func (n *Nasabah) SetAge(tmp []string) {
	n.Age, err = strconv.Atoi(tmp[2])
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
}

func (n *Nasabah) SetBalance(tmp []string) {
	balanced, err := strconv.ParseFloat(tmp[3], 64)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	n.Balanced = balanced
}

func (n *Nasabah) SetPreviousBalance(tmp []string) {
	prevbalanced, err := strconv.ParseFloat(tmp[4], 64)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	n.PreviousBalanced = prevbalanced
}

func (n *Nasabah) SetFreeTransfer(tmp []string) {
	transfer, err := strconv.ParseFloat(tmp[6], 64)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	n.FreeTransfer = transfer
}

func (n *Nasabah) SetNo1() {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	num, _ := strconv.ParseUint(string(b), 10, 64)

	n.No1 = int(num)
}
func (n *Nasabah) SetNo2a() {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	num, _ := strconv.ParseUint(string(b), 10, 64)

	n.No2a = int(num)
}
func (n *Nasabah) SetNo2b() {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	num, _ := strconv.ParseUint(string(b), 10, 64)

	n.No2b = int(num)
}
func (n *Nasabah) SetNo3(a int) {
	// b := make([]byte, 64)
	// b = b[:runtime.Stack(b, false)]
	// b = bytes.TrimPrefix(b, []byte("goroutine "))
	// b = b[:bytes.IndexByte(b, ' ')]
	// num, _ := strconv.ParseUint(string(b), 10, 64)

	n.No3 = a
}

func (n *Nasabah) Setter(tmp []string) {
	n.SetId(tmp)
	n.SetName(tmp)
	n.SetAge(tmp)
	n.SetBalance(tmp)
	n.SetPreviousBalance(tmp)
	n.SetFreeTransfer(tmp)
}

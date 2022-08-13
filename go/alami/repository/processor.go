package repository

import (
	"alami/model"
)

type Process interface {
	GetBankData(path string) []*model.Nasabah
	CreateCSV(nasabahList []*model.Nasabah)
}

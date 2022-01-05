package main

import (
	"github.com/anil8753/onesheds/apps/warehousemen/service/db"
	"github.com/anil8753/onesheds/apps/warehousemen/service/interfaces"
	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
)

type HandlerDependency struct {
	DBInst interfaces.Database
	Ledger *ledger.Ledger
}

func (i *HandlerDependency) GetDB() interfaces.Database {
	return i.DBInst
}

func (i *HandlerDependency) GetLedger() *ledger.Ledger {
	return i.Ledger
}

func NewHandlerDependency() *HandlerDependency {

	h := HandlerDependency{
		Ledger: &ledger.Ledger{},
		DBInst: db.NewLevelDB("generic"),
	}

	h.Ledger.Init()

	return &h
}

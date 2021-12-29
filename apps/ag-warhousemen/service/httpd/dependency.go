package main

import (
	"github.com/anil8753/onesheds/apps/warehousemen/service/db"
	"github.com/anil8753/onesheds/apps/warehousemen/service/interfaces"
)

type HandlerDependency struct {
	DBInst interfaces.Database
}

func (i *HandlerDependency) GetDB() interfaces.Database {
	return i.DBInst
}

func NewHandlerDependency() *HandlerDependency {

	h := HandlerDependency{
		DBInst: db.New(),
	}

	return &h
}

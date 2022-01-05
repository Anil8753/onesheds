package interfaces

import "github.com/anil8753/onesheds/apps/warehousemen/service/ledger"

type HandlerDependency interface {
	GetDB() Database
	GetLedger() *ledger.Ledger
}

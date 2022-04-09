package interfaces

import "github.com/anil8753/onesheds/ledger"

type HandlerDependency interface {
	GetDB() Database
	GetLedger() *ledger.Ledger
}

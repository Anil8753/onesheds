package warehouse

import (
	"github.com/anil8753/onesheds/apps/warehousemen/service/interfaces"
	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
)

type Asset struct {
	Database interfaces.Database
	Ledger   *ledger.Ledger
}

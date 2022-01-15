package warehouse

import (
	"github.com/anil8753/onesheds/apps/warehousemen/service/interfaces"
	"github.com/anil8753/onesheds/apps/warehousemen/service/ledger"
)

type Warehouse struct {
	Database interfaces.Database
	Ledger   *ledger.Ledger
}

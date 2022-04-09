package review

import (
	"github.com/anil8753/onesheds/interfaces"
	"github.com/anil8753/onesheds/ledger"
)

type Handler struct {
	Database interfaces.Database
	Ledger   *ledger.Ledger
}

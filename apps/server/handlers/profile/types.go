package profile

import (
	"github.com/anil8753/onesheds/interfaces"
	"github.com/anil8753/onesheds/ledger"
)

type Profile struct {
	Database interfaces.Database
	Ledger   *ledger.Ledger
}

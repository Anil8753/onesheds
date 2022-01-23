package txn

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func RejectOrder(ctx contractapi.TransactionContextInterface) error {
	return ChangeOrderStatusByWarehouseOwner(ctx, OrderStatusRejected)
}

package txn

import (
	"errors"
	"fmt"

	"github.com/anil8753/onesheds/chaincode/core/asset"
	"github.com/anil8753/onesheds/chaincode/core/utils"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func isReadAllowed(ctx contractapi.TransactionContextInterface, order *Order) error {

	clientMSP, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return err
	}

	// Regulator is always allowed
	if clientMSP == utils.RegulatorMSP {
		return nil
	}

	// Warehouse owner can access
	if err := IsUserWarehouseOwner(ctx, order.WarehouseId); err == nil {
		return nil
	}

	// Depositor can access
	if err := ctx.GetClientIdentity().AssertAttributeValue("userId", order.DepositorId); err == nil {
		return nil
	}

	return errors.New("only regulator or warehouseman or depositor can access")
}

func IsUserWarehouseOwner(ctx contractapi.TransactionContextInterface, warehouseId string) error {

	// Warehouse owner can access
	wh, err := asset.Query(ctx, warehouseId)
	if err != nil {
		return fmt.Errorf("asset.Query failed. %w", err)
	}

	cid := ctx.GetClientIdentity()
	if err := cid.AssertAttributeValue("userId", wh.OwnerId); err != nil {
		return err
	}

	return nil
}

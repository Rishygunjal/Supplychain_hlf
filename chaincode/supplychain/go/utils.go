
package main

import (
    "encoding/json"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Helper function to convert the structure to JSON and put it in world state
func putState(ctx contractapi.TransactionContextInterface, key string, data interface{}) error {
    dataAsBytes, err := json.Marshal(data)
    if err != nil {
        return err
    }

    return ctx.GetStub().PutState(key, dataAsBytes)
}

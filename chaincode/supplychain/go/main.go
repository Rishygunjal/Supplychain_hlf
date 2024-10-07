package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract handles the lifecycle operations of products
type SmartContract struct {
    contractapi.Contract
}

// Product represents the details of a product in the supply chain system
type Product struct {
    ProductID          string `json:"productID"`
    Name               string `json:"name"`
    Description        string `json:"description"`
    ManufacturingDate  string `json:"manufacturingDate"`
    BatchNumber        string `json:"batchNumber"`
    Status             string `json:"status"`
    SupplyDate         string `json:"supplyDate"`
    WarehouseLocation  string `json:"warehouseLocation"`
    WholesaleDate      string `json:"wholesaleDate"`
    WholesaleLocation  string `json:"wholesaleLocation"`
    Quantity           int    `json:"quantity"`
}

// SeedLedger populates the ledger with initial product records
func (sc *SmartContract) SeedLedger(ctx contractapi.TransactionContextInterface) error {
    initialProducts := []Product{
        {
            ProductID: "PRD001", 
            Name: "Widget A", 
            Description: "Basic Widget A Description", 
            ManufacturingDate: "2023-09-25", 
            BatchNumber: "BATCH001", 
            Status: "Manufactured",
        },
        {
            ProductID: "PRD002", 
            Name: "Widget B", 
            Description: "Basic Widget B Description", 
            ManufacturingDate: "2023-09-26", 
            BatchNumber: "BATCH002", 
            Status: "Manufactured",
        },
    }

    for _, product := range initialProducts {
        productJSON, err := json.Marshal(product)
        if err != nil {
            return fmt.Errorf("Failed to serialize product data: %v", err)
        }

        err = ctx.GetStub().PutState(product.ProductID, productJSON)
        if err != nil {
            return fmt.Errorf("Error adding product %s to ledger: %v", product.ProductID, err)
        }
    }
    return nil
}

func main() {
    sc := new(SmartContract)
    chaincode, err := contractapi.NewChaincode(sc)
    if err != nil {
        fmt.Printf("Error creating supply chain chaincode: %v\n", err)
        return
    }

    if err := chaincode.Start(); err != nil {
        fmt.Printf("Failed to start chaincode: %v\n", err)
    }
}

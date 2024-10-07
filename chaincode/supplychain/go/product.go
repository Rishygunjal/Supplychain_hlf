package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// RegisterProduct adds a new product record to the ledger
func (sc *SmartContract) RegisterProduct(ctx contractapi.TransactionContextInterface, id string, name string, details string, mfgDate string, batch string) error {
    newProduct := Product{
        ProductID:       id,
        Name:            name,
        Description:     details,
        ManufacturingDate: mfgDate,
        BatchNumber:     batch,
        Status:          "Registered",
    }

    productBytes, err := json.Marshal(newProduct)
    if err != nil {
        return fmt.Errorf("unable to marshal product: %v", err)
    }

    return ctx.GetStub().PutState(id, productBytes)
}

// UpdateSupplyInfo modifies the product status with supply details
func (sc *SmartContract) UpdateSupplyInfo(ctx contractapi.TransactionContextInterface, id string, supplyDate string, location string) error {
    product, err := sc.GetProductByID(ctx, id)
    if err != nil {
        return err
    }

    product.SupplyDate = supplyDate
    product.WarehouseLocation = location
    product.Status = "In Supply"

    updatedBytes, err := json.Marshal(product)
    if err != nil {
        return fmt.Errorf("unable to marshal updated product: %v", err)
    }

    return ctx.GetStub().PutState(id, updatedBytes)
}

// RecordWholesale updates the product with wholesale transaction details
func (sc *SmartContract) RecordWholesale(ctx contractapi.TransactionContextInterface, id string, wholesaleDate string, location string, quantity int) error {
    product, err := sc.GetProductByID(ctx, id)
    if err != nil {
        return err
    }

    product.WholesaleDate = wholesaleDate
    product.WholesaleLocation = location
    product.Quantity = quantity
    product.Status = "Wholesale Completed"

    wholesaleBytes, err := json.Marshal(product)
    if err != nil {
        return fmt.Errorf("unable to marshal wholesale details: %v", err)
    }

    return ctx.GetStub().PutState(id, wholesaleBytes)
}

// GetProductByID retrieves a product from the ledger using its ID
func (sc *SmartContract) GetProductByID(ctx contractapi.TransactionContextInterface, id string) (*Product, error) {
    data, err := ctx.GetStub().GetState(id)
    if err != nil {
        return nil, fmt.Errorf("error reading from ledger: %v", err)
    }
    if data == nil {
        return nil, fmt.Errorf("product with ID %s not found", id)
    }

    var product Product
    if err := json.Unmarshal(data, &product); err != nil {
        return nil, fmt.Errorf("error unmarshaling product data: %v", err)
    }

    return &product, nil
}

// ChangeProductStatus updates the current status of a product (e.g., "Sold")
func (sc *SmartContract) ChangeProductStatus(ctx contractapi.TransactionContextInterface, id string, newStatus string) error {
    product, err := sc.GetProductByID(ctx, id)
    if err != nil {
        return err
    }

    product.Status = newStatus
    statusBytes, err := json.Marshal(product)
    if err != nil {
        return fmt.Errorf("error marshaling status update: %v", err)
    }

    return ctx.GetStub().PutState(id, statusBytes)
}

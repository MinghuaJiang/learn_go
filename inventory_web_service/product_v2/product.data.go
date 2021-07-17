package product_v2

import (
	"database/sql"

	"github.com/minghuajiang/learn_go/inventory_web_service/database"
)

func getProduct(productID int) (*Product, error) {
	row := database.DBConn.QueryRow(`SELECT productId,
	manufacturer,
	sku,
	upc,
	pricePerUnit,
	quantityOnHand,
	productName
	FROM products
	WHERE productId = ?`, productID)

	product := &Product{}

	err := row.Scan(&product.ProductID,
		&product.Manufacturer,
		&product.Sku,
		&product.Upc,
		&product.PricePerUnit,
		&product.QuantityOnHand,
		&product.ProductName)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return product, nil
}

func removeProduct(productId int) error {
	_, err := database.DBConn.Exec(`DELETE FROM products where productId = ?`, productId)
	return err
}

func getProductList() ([]Product, error) {
	results, err := database.DBConn.Query(`SELECT productId,
	manufacturer,
	sku,
	upc,
	pricePerUnit,
	quantityOnHand,
	productName
	FROM products`)

	if err != nil {
		return nil, err
	}

	defer results.Close()

	products := make([]Product, 0)

	for results.Next() {
		var product Product
		results.Scan(&product.ProductID,
			&product.Manufacturer,
			&product.Sku,
			&product.Upc,
			&product.PricePerUnit,
			&product.QuantityOnHand,
			&product.ProductName)
		products = append(products, product)
	}

	return products, nil
}

func updateProduct(product Product) error {
	_, err := database.DBConn.Exec(`Update products SET 
	manufacturer=?,
	sku=?,
	upc=?,
	pricePerUnit=CAST(? AS DECIMAL(13,2)),
	quantityOnHand=?,
	productName=?
	WHERE productId=?`,
		product.Manufacturer,
		product.Sku,
		product.Upc,
		product.PricePerUnit,
		product.QuantityOnHand,
		product.ProductName,
		product.ProductID)

	return err
}

func insertProduct(product Product) (int, error) {
	result, err := database.DBConn.Exec(`INSERT INTO products 
	(manufacturer,
		sku,
		upc,
		pricePerUnit,
		quantityOnHand,
		productName) VALUES (?, ?, ?, ?, ?, ?)`,
		product.Manufacturer,
		product.Sku,
		product.Upc,
		product.PricePerUnit,
		product.QuantityOnHand,
		product.ProductName)

	if err != nil {
		return 0, err
	}

	insertID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(insertID), nil
}

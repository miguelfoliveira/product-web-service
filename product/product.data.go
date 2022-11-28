package product

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/miguelfoliveira/go-web-service/database"
)

const queryMaxTime = 15 * time.Second

func insertProduct(product Product) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryMaxTime)
	defer cancel()

	result, err := database.DbConn.ExecContext(ctx, `INSERT INTO products  
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
		product.Name)

	if err != nil {
		log.Println(err.Error())
		return 0, nil
	}

	insertId, err := result.LastInsertId()

	if err != nil {
		log.Println(err.Error())
		return 0, nil
	}

	return int(insertId), nil
}

func deleteProduct(productId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), queryMaxTime)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `DELETE FROM products WHERE productId =?`,
		productId)

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func updateProduct(product Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), queryMaxTime)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `UPDATE products SET 
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
		product.Name,
		product.Id)

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func getProduct(productID int) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryMaxTime)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx,
		`SELECT productId,
		manufacturer,
		sku,
		upc,
		pricePerUnit,
		quantityOnHand,
		productName
		FROM products
		WHERE productId=  ?`, productID)

	product := &Product{}

	err := row.Scan(
		&product.Id,
		&product.Manufacturer,
		&product.Sku,
		&product.Upc,
		&product.PricePerUnit,
		&product.QuantityOnHand,
		&product.Name,
	)

	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return product, nil
}

func getProducts() ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryMaxTime)
	defer cancel()
	rowsResult, err := database.DbConn.QueryContext(ctx, `
		SELECT productId,
		manufacturer,
		sku,
		upc,
		pricePerUnit,
		quantityOnHand,
		productName
		FROM products
	`)

	if err != nil {
		return nil, err
	}

	defer rowsResult.Close()
	products := make([]Product, 0)

	for rowsResult.Next() {
		var product Product
		rowsResult.Scan(
			&product.Id,
			&product.Manufacturer,
			&product.Sku,
			&product.Upc,
			&product.PricePerUnit,
			&product.QuantityOnHand,
			&product.Name,
		)
		products = append(products, product)
	}

	return products, nil
}

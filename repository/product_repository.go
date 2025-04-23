package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT * FROM product"

	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {

	var id int

	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		"VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {
	var product model.Product

	// prepare a query
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE ID = $1")

	// verify if exists erro in prepare
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Execute query and scan copies the columns from the matched row into the values
	// pointed at by dest &product.ID (endereço de memória do product.ID)
	err = query.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		fmt.Println(err)

		// Check if no rows were found
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	// Close query
	query.Close()

	return &product, nil
}

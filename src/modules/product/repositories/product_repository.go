package repositories

import (
	"database/sql"

	"github.com/henrique77/product-api/src/modules/product/entities"
)

type (
	ProductRepository interface {
		Get() ([]*entities.Products, error)
		GetOne(int) (*entities.Products, error)
		Create(*entities.Products) error
		Update(int, *entities.Products) (*entities.Products, error)
		Delete(int) error
	}

	productRepository struct {
		*sql.DB
	}
)

func (p *productRepository) Create(product *entities.Products) error {

	_, err := p.DB.Exec("INSERT INTO products (name, price, description, qtd) VALUES (?, ?, ?, ?)", product.Name, product.Price, product.Description, product.Qtd)
	if err != nil {
		return err
	}

	return nil
}

func (p *productRepository) Get() ([]*entities.Products, error) {
	rows, err := p.DB.Query("SELECT product_id, name, price, description, qtd from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entities.Products
	for rows.Next() {
		var product entities.Products
		if err := rows.Scan(&product.ProductId, &product.Name, &product.Price, &product.Description, &product.Qtd); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (p *productRepository) GetOne(id int) (*entities.Products, error) {
	var product entities.Products
	err := p.DB.QueryRow("SELECT product_id, name, price, description, qtd from products WHERE product_id = ?", id).
		Scan(&product.ProductId, &product.Name, &product.Price, &product.Description, &product.Qtd)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepository) Update(id int, product *entities.Products) (*entities.Products, error) {
	_, err := p.DB.Exec("UPDATE products SET name = ?, price = ?, description = ?, qtd = ? WHERE product_id = ? ",
		&product.Name, &product.Price, &product.Description, &product.Qtd, &id)
	if err != nil {
		return nil, err
	}

	productNew, err := p.GetOne(id)
	if err != nil {
		return nil, err
	}

	return productNew, nil
}

func (p *productRepository) Delete(id int) error {
	result, err := p.Exec("DELETE FROM products WHERE product_id = ?", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

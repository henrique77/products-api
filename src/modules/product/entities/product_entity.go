package entities

type Products struct {
	ProductId   int32   `db:"product_id" json:"id"`
	Name        string  `db:"name" json:"name"`
	Price       float32 `db:"price" json:"price"`
	Description string  `db:"description" json:"description"`
	Qtd         int32   `db:"qtd" json:"qtd"`
}

type ProductsList []*Products

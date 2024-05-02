USE `product-api`;

CREATE TABLE products (
    product_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    price FLOAT,
    description TEXT,
    qtd INT
);
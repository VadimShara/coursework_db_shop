CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE products (
    product_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name_product VARCHAR(50) NOT NULL,
    product_unit VARCHAR(20),
    purchase_price REAL NOT NULL,
    selling_price REAL NOT NULL
);

CREATE TABLE sellers (
    seller_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name_seller VARCHAR(50) NOT NULL, 
    commission_rate REAL NOT NULL
);

CREATE TABLE sales (
    sale_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL,
    seller_id UUID NOT NULL,
    number INTEGER NOT NULL,
    saleDateTime TIMESTAMP NOT NULL, 
    FOREIGN KEY (product_id) REFERENCES products(product_id),
    FOREIGN KEY (seller_id) REFERENCES sellers(seller_id)
)
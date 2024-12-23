package models

type Product struct {
	Product_id int
    Name_product string
    Product_unit string
    Purchase_price float32
    Selling_price float32
}

type Sellers struct {
	Seller_id int
    Name_seller string
    Commission_rate float32
}

type Sales struct {
	Sale_id int
    Product_id int
    Seller_id int
    Number int
    SaleDateTime int
}
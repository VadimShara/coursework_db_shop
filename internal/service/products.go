package service

import (
	"fmt"

	"github.com/VadimShara/coursework_db_shop/internal/models"
	"github.com/VadimShara/coursework_db_shop/internal/repo"
)

func (s *Service) AddProduct() {
	var p models.Product

	fmt.Println("Введите наименование товара:")
	fmt.Scanln(&p.name_product)

	fmt.Println("Введите единица измерения товара:")
	fmt.Scanln(&p.product_unit)

	fmt.Println("Введите цену закупки:")
	fmt.Scanln(&p.purchase_price)

	fmt.Println("Введите цену продажи:")
	fmt.Scanln(&p.selling_price)

	s.db.AddProduct(p)
}

func (s *Service) RemoveProducts() {
	var p models.Product

	fmt.Println("Введите id товара:")
	fmt.Scanln(&p.product_id)

	s.db.DeleteProduct(p.product_id)
}

func (s *Service) ListProducts() {
	products, _ := s.db.ListProducts()

	fmt.Println(products)
}

func (s *Service) SearchProduct() {
	var p models.Product

	fmt.Println("Введите наименование товара:")
	fmt.Scanln(&p.name_product)

	p,_ = s.db.SearchProduct(p.name_product)
	fmt.Println(p)
}


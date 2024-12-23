package service

import (
	"fmt"

	"github.com/VadimShara/coursework_db_shop/internal/models"
)

func (s *Service) AddSale() {
	var sale models.Sale

	fmt.Println("Введите фамилию и инициалы сотрудника:")
	fmt.Scanln(&sale.name_seller)

	fmt.Println("Введите процент коммисионных сотрудника:")
	fmt.Scanln(&sale.commission_rate)

	
}

func (s *Service) RemoveSale() {
	var seller models.Seller

	fmt.Println("Введите id сотрудника:")
	fmt.Scanln(&seller.seller_id)

}

func ListSales() {

}

// func (s *Service) SearchSale() error {
// 	var sale models.Sale

// 	fmt.Println("Введите фамилию и инициалы сотрудника:")
// 	fmt.Scanln(&s.name_seller)
// }
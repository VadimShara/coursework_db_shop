package service

import (
	"fmt"

	"github.com/VadimShara/coursework_DB_shop/internal/models"
)

func (s *Service) AddSeller() {
	var seller models.Seller

	fmt.Println("Введите фамилию и инициалы сотрудника:")
	fmt.Scanln(&seller.name_seller)

	fmt.Println("Введите процент коммисионных сотрудника:")
	fmt.Scanln(&seller.commission_rate)

	s.db.AddSeller(seller)
}

func (s *Service) RemoveSeller() {
	var seller models.Seller

	fmt.Println("Введите id сотрудника:")
	fmt.Scanln(&seller.seller_id)

	s.db.DeleteSeller(seller.seller_id)
}

func (s *Service) ListSellers() {
	sellers, _ := s.db.ListSellers
	fmt.Println(sellers)
}

func (s *Service) SearchSeller() {
	var seller models.Seller

	fmt.Println("Введите фамилию и инициалы сотрудника:")
	fmt.Scanln(&seller.name_seller)

	seller, _ = s.db.SearchSeller(seller)
	fmt.Println(seller)
}
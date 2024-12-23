package service

import(
	"fmt"
)

func admin() {
	var section int
	var action int

	fmt.Println("Выберите раздел:")
	fmt.Println("1 - сотрудники")
	fmt.Println("2 - товары")
	fmt.Println("3 - продажи")
	fmt.Scanln(&section)

	fmt.Println("Выберите действие:")
	fmt.Println("1 - добавить")
	fmt.Println("2 - удалить")
	fmt.Println("3 - вывести список")
	fmt.Println("4 - найти")
	fmt.Scanln(&action)

	switch section{
	case 1:
		switch action {
		case 1:

		case 2:

		case 3:	

		}
	case 2:
		switch action {
		case 1:

		case 2:

		case 3:	

		}
	case 3:
		switch action {
		case 1:

		case 2:

		case 3:	

		}			
	}

}
package service

import(
    "fmt"
)

func Auth() {
    var role int
    fmt.Println("Укажите вашу роль:\n0 - admin\n1 - сотрудник")
    fmt.Scanln(&role)

    switch role {
    case 0:
        fmt.Println("ADMIN")
		admin()
    case 1:
        fmt.Println("EMPLOYEE") 
		emloyee()
    default:
        fmt.Println("ENTRY ERROR")   
        Auth() 
    }

}
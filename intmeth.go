package main

import "fmt"

type Account struct {
	Total_balance int
	Amount        int
}

func (a *Account) Add() int {
	fmt.Println("Enter deposit amount")
	fmt.Scan(&a.Amount)
	a.Total_balance = a.Total_balance + a.Amount
	return a.Total_balance
}

func (a *Account) withdraw() int {
	fmt.Println("Enter withdraw amount")
	fmt.Scan(&a.Amount)
	return a.Total_balance - a.Amount

}

type Account_interface interface {
	Add()
	withdraw()
}

func main() {
	var choice int
	var count int = 0
	var user = Account{Total_balance: 50}
	//fmt.Print("Account Holder Details:", user)
	fmt.Print("For add Money in your Account Enter 1, for withdraw money enter: 2 and for Exit: 0= ")
	fmt.Scan(&choice)
	for count == 0 {
		if choice == 1 {
			fmt.Println("Current balance is:", user.Add())
			fmt.Print("For add Money in your Account Enter 1, for withdraw money Enter: 2 and for Exit: 0= ")
			fmt.Scan(&choice)

		} else if choice == 2 {
			fmt.Println("Current balance is:", user.withdraw())
			fmt.Print("For add Money in your Account Enter 1, for withdraw money Enter: 2 and for Exit: 0= ")
			fmt.Scan(&choice)
		} else {
			count = 1

		}
	}

}

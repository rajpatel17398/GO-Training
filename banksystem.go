package main

import "fmt"

// type para struct {
// 	amount int
// }

func deposit(am int) {
	amount1 := 100
	fmt.Println("your ammount is 100 initially")
	depo := am + amount1
	fmt.Println("your current ammount is:", depo)
}
func Withdraw(am int) {
	amount1 := 100
	fmt.Println("your ammount is 100 initially")
	depo := amount1 - am
	fmt.Println("your current ammount is:", depo)
	if depo < 0 {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("your balance is", depo, "Please add the money")
			}
		}()
		panic("your ammount is not in minus")
	}
}

func main() {
	fmt.Println("welcome to the bank system")
	fmt.Println("your ammount is 100 initially")

	fmt.Println("Deposit or Withdraw")
	var dw string
	fmt.Scanln(&dw)
	var am int
	fmt.Println("Enter the ammount")
	fmt.Scanln(&am)

	if dw == "Deposit" || dw == "deposit" {
		deposit(am)
	} else if dw == "Withdraw" || dw == "withdraw" {
		Withdraw(am)
	}

}

package main

import "fmt"

func main() {

}

type PaymentProcessor interface {
	Pay(amount int) error
}

type CreditUnion struct{}

func (c CreditUnion) Pay(amount int) error {
	fmt.Printf("Paid %d using credit card", amount)
	return nil
}

type Upi struct{}

func (u Upi) Pay(amount int) error {
	fmt.Printf("Paid %d amount using UPI", amount)
	return nil
}

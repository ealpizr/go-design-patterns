package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Paying using cash")
}

type CardPayment struct{}

func (CardPayment) Pay(cardNumber int) {
	fmt.Printf("Paying using card %d\n", cardNumber)
}

type CardPaymentAdapter struct {
	cardPayment *CardPayment
	cardNumber  int
}

func (cpa CardPaymentAdapter) Pay() {
	cpa.cardPayment.Pay(cpa.cardNumber)
}

func ProcessPayment(p Payment) {
	p.Pay()
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	card := &CardPaymentAdapter{
		cardPayment: &CardPayment{},
		cardNumber:  123456789,
	}
	ProcessPayment(card)
}

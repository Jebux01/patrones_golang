package main

import (
	"fmt"
	"math/rand"
)

type Payment interface {
	Pay(params map[string]interface{}) (string, bool)
}

type CashPayment struct{}

func (c *CashPayment) Pay(params map[string]interface{}) (string, bool) {
	amount, ok := params["amount"].(float32)
	if !ok {
		panic("Amount not found")
	}

	discount, ok := params["discount"].(float32)
	if !ok {
		discount = 0
	}

	if randomNum() {
		discountedAmount := amount * (1 - discount/100) // Aplica descuento
		formattedString := fmt.Sprintf("Payment with cash of %.2f was successful\n", discountedAmount)
		return formattedString, true
	}

	return "Payment with cash failed", false
}

type CardPayment struct{}

func (c *CardPayment) Pay(params map[string]interface{}) (string, bool) {
	amount, ok := params["amount"].(float32)
	if !ok {
		panic("Amount not found")
	}
	discount, ok := params["discount"].(float32)
	if !ok {
		discount = 0
	}

	if randomNum() {
		discountedAmount := amount * (1 - discount/100) // Aplica descuento
		formattedString := fmt.Sprintf("Payment with Card of %.2f was successful\n", discountedAmount)
		return formattedString, true
	}

	return "Payment with Card failed", false
}

type PSEPayment struct{}

func (c *PSEPayment) Pay(params map[string]interface{}) (string, bool) {
	amount, ok := params["amount"].(float32)
	if !ok {
		panic("Amount not found")
	}
	discount, ok := params["discount"].(float32)
	if !ok {
		discount = 0
	}

	if randomNum() {
		discountedAmount := amount * (1 - discount/100) // Aplica descuento
		formattedString := fmt.Sprintf("Payment with PSE of %.2f was successful\n", discountedAmount)
		return formattedString, true
	}

	return "Payment with PSE failed", false
}

type PaymentDecorator struct {
	Payment
	discount float32
}

func NewPaymentDecorator(p Payment, discount float32) *PaymentDecorator {
	return &PaymentDecorator{Payment: p, discount: discount}
}

func (d *PaymentDecorator) Pay(params map[string]interface{}) (string, bool) {
	params["discount"] = d.discount
	return d.Payment.Pay(params)
}

func randomNum() bool {
	numRandom := rand.Intn(100)
	if numRandom%3 == 0 || numRandom%5 == 0 {
		return true
	}
	if numRandom%2 == 0 {
		return true
	}
	return false
}

package main

import (
	"fmt"
	"math/rand"
)

type Payment interface {
	Pay(params map[string]interface{}) (string, bool)
	Undo(params map[string]interface{}) (string, bool)
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

func (c *CashPayment) Undo(params map[string]interface{}) (string, bool) {
	return "Payment undo", true
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

func (c *CardPayment) Undo(params map[string]interface{}) (string, bool) {
	return "Payment undo", true
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

func (c *PSEPayment) Undo(params map[string]interface{}) (string, bool) {
	return "Payment undo", true
}

type PaymentDecorator struct {
	Payment
	discount float32
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

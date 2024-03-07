package main

import (
	"fmt"

	"github.com/google/uuid"
)

type PaymentObserver interface {
	Update(string, bool, string)
}

type PaymentLogger struct{}

func (pl *PaymentLogger) Update(method string, success bool, message string) {
	if success {
		fmt.Printf(message)
		return
	}

	fmt.Printf("Payment with %s failed\n", method)
}

type PaymentProcessor struct {
	paymentStrategies map[string]Payment
	observers         []PaymentObserver
}

func NewPaymentProcessor() *PaymentProcessor {
	return &PaymentProcessor{
		paymentStrategies: make(map[string]Payment),
		observers:         make([]PaymentObserver, 0),
	}
}

func (p *PaymentProcessor) RegisterPaymentMethod(name string, payment Payment) {
	p.paymentStrategies[name] = payment
}

func (p *PaymentProcessor) Attach(observer PaymentObserver) {
	p.observers = append(p.observers, observer)
}

func (p *PaymentProcessor) GetPaymentMethod(method string) (Payment, bool) {
	paymentCommand, ok := p.paymentStrategies[method]
	return paymentCommand, ok
}

func savePayment(success bool, method string, amount float32) string {
	id := uuid.New()
	data := map[string]interface{}{
		"uuid":    id.String(),
		"amount":  amount,
		"method":  method,
		"success": success,
	}

	saveReg(id.String(), fmt.Sprint(data))
	return id.String()
}

package main

import "fmt"

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

func (p *PaymentProcessor) Pay(params map[string]interface{}) bool {
	payment, ok := p.paymentStrategies[params["name"].(string)]
	if !ok {
		panic("Payment method not found")
	}
	message, success := payment.Pay(params)
	for _, observer := range p.observers {
		observer.Update(params["name"].(string), success, message)
	}
	return success
}

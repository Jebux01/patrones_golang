package main

func main() {
	paymentProcessor := NewPaymentProcessor()

	cashPayment := &CashPayment{}
	cashPaymentWithDiscount := NewPaymentDecorator(cashPayment, 10.0)
	paymentProcessor.RegisterPaymentMethod("cash", cashPaymentWithDiscount)

	cardPayment := &CardPayment{}
	cardPaymentWithDiscount := NewPaymentDecorator(cardPayment, 5.0)
	paymentProcessor.RegisterPaymentMethod("card", cardPaymentWithDiscount)

	psePayment := &PSEPayment{}
	psePaymentWithDiscount := NewPaymentDecorator(psePayment, 15.0)
	paymentProcessor.RegisterPaymentMethod("pse", psePaymentWithDiscount)

	paymentLogger := &PaymentLogger{}
	paymentProcessor.Attach(paymentLogger)

	paymentProcessor.Pay(map[string]interface{}{"name": "cash", "amount": float32(100.00)})
	paymentProcessor.Pay(map[string]interface{}{"name": "card", "amount": float32(100.00)})
	paymentProcessor.Pay(map[string]interface{}{"name": "pse", "amount": float32(100.00)})
}

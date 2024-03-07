package main

func NewPaymentDecorator(p Payment, discount float32) *PaymentDecorator {
	return &PaymentDecorator{Payment: p, discount: discount}
}

func (d *PaymentDecorator) Pay(params map[string]interface{}) (string, bool) {
	params["discount"] = d.discount
	return d.Payment.Pay(params)
}

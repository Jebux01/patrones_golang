package main

func NotifyObservers(observers []PaymentObserver, method string, success bool, message string) {
	for _, observer := range observers {
		observer.Update(method, success, message)
	}
}

func (p *PaymentProcessor) Execute(payment Payment, params map[string]interface{}) bool {
	message, success := payment.Pay(params)
	uuid := savePayment(success, params["name"].(string), params["amount"].(float32))
	message = message + " with id: " + uuid
	NotifyObservers(p.observers, params["name"].(string), success, message)
	return success
}

func (p *PaymentProcessor) Detach(payment Payment, params map[string]interface{}) bool {
	message, success := payment.Undo(params)
	delReg(params["uuid"].(string))
	NotifyObservers(p.observers, params["name"].(string), success, message)
	return success
}

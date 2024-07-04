package main

import "fmt"

type Payment interface {
	processPayment(amount float32) string
}

type CreditCardPayment struct {
	cardNumber string
}

func (c CreditCardPayment) processPayment(amount float32) string {
	return fmt.Sprintf("Processed payment of %f through credit card %s", amount, c.cardNumber)
}

type DebitCardPayment struct {
	cardNumber string
}

func (d DebitCardPayment) processPayment(amount float32) string {
	return fmt.Sprintf("Processed payment of %f through debit card %s", amount, d.cardNumber)
}

type PaypalPayment struct {
	email string
}

func (p PaypalPayment) processPayment(amount float32) string {
	return fmt.Sprintf("Processed payment of %f through PayPal account %s", amount, p.email)
}

type PaymentFactory struct{}

func (f PaymentFactory) CreatePayment(paymentType string, details interface{}) Payment {
	switch paymentType {
	case "creditcard":
		if cardNumber, ok := details.(string); ok {
			return CreditCardPayment{cardNumber: cardNumber}
		}
	case "debitcard":
		if cardNumber, ok := details.(string); ok {
			return DebitCardPayment{cardNumber: cardNumber}
		}
	case "paypal":
		if email, ok := details.(string); ok {
			return PaypalPayment{email: email}
		}
	default:
		panic("Unsupported payment type")
	}
	return nil
}

func main() {
	factory := PaymentFactory{}

	creditCardPayment := factory.CreatePayment("creditcard", "1234-5678-9012-3456")
	fmt.Println(creditCardPayment.processPayment(100.00))

	debitCardPayment := factory.CreatePayment("debitcard", "2345-6789-0123-4567")
	fmt.Println(debitCardPayment.processPayment(200.00))

	paypalPayment := factory.CreatePayment("paypal", "user@example.com")
	fmt.Println(paypalPayment.processPayment(300.00))
}

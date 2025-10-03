package main

import (
	"fmt"
)

/*
Применяется для интеграции старого кода с новым, либо же с внешним кодом
в исходники которого мы не можем залезть.
Позволяет клиенту работать с единым интерфейсом
Код клиента не знает о деталях адаптации, поэтому можно добавлять новые адаптеры
без изменения клиентского кода.
Из минусов можно выделить - 1) усложнение кода, 2) небольшой оверхед в рантайме
на типизацию

Реальные примеры:
1) http.HandlerFunc — адаптирует обычную функцию func(w http.ResponseWriter, r *http.Request) к интерфейсу http.Handler
2) database/sql драйверы — внешние пакеты реализуют набор интерфейсов driver.Conn, driver.Stmt, и «адаптируются» к унифицированному API database/sql
3) Логгеры: адаптеры между log.Logger и сторонними логгерами (например, zap.NewStdLog делает адаптацию Zap к stdlib-логгеру
*/

type PayPal struct{}

func (p *PayPal) MakePayment(amount float64) string {
	return fmt.Sprintf("PayPal: оплата $%.2f выполнена", amount)
}

type Stripe struct{}

func (s *Stripe) Charge(amount float64) string {
	return fmt.Sprintf("Stripe: списано $%.2f", amount)
}

type PaymentProcessor interface {
	ProcessPayment(amount float64) string
}

type PayPalAdapter struct {
	paypal *PayPal
}

func NewPayPalAdapter() PaymentProcessor {
	return &PayPalAdapter{paypal: &PayPal{}}
}

func (p *PayPalAdapter) ProcessPayment(amount float64) string {
	return p.paypal.MakePayment(amount)
}

type StripeAdapter struct {
	stripe *Stripe
}

func NewStripeAdapter() PaymentProcessor {
	return &StripeAdapter{stripe: &Stripe{}}
}

func (s *StripeAdapter) ProcessPayment(amount float64) string {
	return s.stripe.Charge(amount)
}

func demonstratePaymentAdapter() {
	processors := []PaymentProcessor{
		NewPayPalAdapter(),
		NewStripeAdapter(),
	}

	for _, processor := range processors {
		result := processor.ProcessPayment(99.99)
		fmt.Println(result)
	}
}

func main() {
	demonstratePaymentAdapter()
}

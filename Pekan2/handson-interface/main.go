package main

import "fmt"

type PaymentMethod interface {
	Pay() string
}

type CreditCard struct{}

func (c CreditCard) Pay() string {
	return "Pembayaran dengan Kartu Kredti Berhasil"
}

type PayPal struct{}

func (p PayPal) Pay() string {
	return "Pembayaran dengan PayPal Berhasil"
}

type BankTransfer struct{}

func (b BankTransfer) Pay() string {
	return "Pembayaran dengan Bank Transfer Berhasil"
}

func ProcessPayment(input PaymentMethod) {
	input.Pay()
}

func main() {
	kartu := CreditCard{}
	payPal := PayPal{}
	tf := BankTransfer{}

	fmt.Println("Hasil Pembayaran:")
	fmt.Println(kartu.Pay())
	fmt.Println(payPal.Pay())
	fmt.Println(tf.Pay())

	ProcessPayment(kartu)
	ProcessPayment(payPal)
	ProcessPayment(tf)
}
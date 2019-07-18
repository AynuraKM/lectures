package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPayByWallet(t *testing.T) {
	amount := 50
	payment := 30
	w := &Wallet{amount}
	w.Pay(payment)

	assert.Equal(t, w.funds, amount-payment)
}

func TestPayByWalletGivesError(t *testing.T) {
	amount := 50
	payment := 60
	w := &Wallet{amount}
	err := w.Pay(payment)
	assert.Error(t, err)
}

func TestBuyByWallet(t *testing.T) {
	amount := 50
	payment := 60
	w := &Wallet{amount}
	_ = Buy(w, payment)
	assert.Equal(t, w.funds, amount-payment)
}

func TestBuyByCreditCard(t *testing.T) {
	amount := 50
	payment := 30
	c := &CreditCard{funds: amount}
	_ = Buy(c, payment)
	assert.Equal(t, c.funds, amount-payment)
}

func TestBuyByBitcoin(t *testing.T) {
	payment := 30
	b := &Bitcoin{}
	b.transactions = make(map[int]int)
	b.transactions[0] = 150
	b.transactions[1] = -10
	b.transactions[2] = 10
	_ = Buy(b, payment)
	fmt.Println(b)
}

func TestCheckAndBuyWallet(t *testing.T) {
	amount := 50
	w := &Wallet{amount}
	CheckAndBuy(w, amount)
	assert.Equal(t, w.funds, amount)
}

func TestCheckPaymentType(t *testing.T) {
	amount := 50
	c := &CreditCard{amount, "Ainura", time.Time{}, nil}
	fmt.Println(CheckPaymentType(c))
}

func TestBitcoin(t *testing.T) {
	b := &Bitcoin{}
	b.transactions = make(map[int]int)
	b.transactions[0] = 150
	b.transactions[1] = -20
	b.transactions[2] = 10
	fmt.Println(b)
}

func TestGetFundsCredit(t *testing.T) {
	w := &Wallet{50}
	fmt.Println(w.GetFunds())
}

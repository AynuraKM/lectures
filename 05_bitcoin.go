package main

import (
	"errors"
	"fmt"
	"time"
)

// имитация платежной системы

type Wallet struct {
	funds int
}

type CreditCard struct {
	funds      int
	owner      string
	expireTime time.Time
	bonuses    []Bonus
}

type Bitcoin struct {
	transactions map[int]int
	SumOfTrans   int
}

type Bonus struct {
	name        string
	description string
}

type Payer interface {
	Pay(int) error
}

type Funder interface {
	GetFunds() int
}

type PayFunder interface {
	Payer
	Funder
}

func (w *Wallet) Pay(amount int) error {
	if amount > w.funds {
		return fmt.Errorf("недостаточно средств")
	}
	w.funds -= amount
	return nil
}

func (c *CreditCard) Pay(amount int) error {
	if amount > c.funds {
		return fmt.Errorf("недостаточно средств")
	}

	c.funds -= amount
	return nil
}

//func (b *Bitcoin) init() {
//	b.transactions = make(map[int]int)
//	fmt.Println(b)
//}

func (b *Bitcoin) Pay(amount int) error {
	SumAllTrans := 0
	if len(b.transactions) == 0 {
		return errors.New("not enough funds")
	}

	for i := 0; i < len(b.transactions); i++ {
		SumAllTrans += b.transactions[i]
	}
	b.SumOfTrans = SumAllTrans
	if SumAllTrans < amount {
		return errors.New("not enough funds")
	}

	b.transactions[len(b.transactions)+1] = amount * (-1)
	return nil
}

func Buy(p Payer, amount int) error {
	err := p.Pay(amount)
	if err != nil {
		return err
	}
	return nil
}

func CheckAndBuy(p PayFunder, amount int) error {
	if p.GetFunds() <= 0 {
		fmt.Println("пополните счет")
		return nil
	}

	err := p.Pay(amount)
	if err != nil {
		return err
	}
	return nil
}

func (w *Wallet) GetFunds() int {
	return w.funds
}

func CheckPaymentType(p Payer) interface{} {
	switch p.(type) {
	case *Wallet:
		fmt.Println("ты пользуешься кошельком")
		return p.(*Wallet).funds
	case *CreditCard:
		fmt.Println("ты пользуешься кредиткой ")
		fmt.Println(p.(*CreditCard).owner)
		return p.(*CreditCard).funds
	case *Bitcoin:
		fmt.Println("Вы используете Биткоин")
		fmt.Println(p.(*Bitcoin).transactions)
		return p.(*Bitcoin).SumOfTrans

	default:
		return nil
	}
}

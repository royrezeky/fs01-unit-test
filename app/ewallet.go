package ewallet

import (
	"errors"
)

func init() {
	isTest = false
}

var saldo float64 = 0
var isTest = true

type Ewallet interface {
	Withdraw(uang float64) (float64, error)
	Deposit(uang float64) (float64, error)
	Print() float64
}

type EwalletImp struct{}

func (e EwalletImp) Withdraw(uang float64) (float64, error) {
	if saldo < uang {
		return 0, errors.New("saldo anda tidak mencukupi")
	}
	saldo = saldo - uang
	return saldo, nil
}

func (e EwalletImp) Deposit(uang float64) (float64, error) {
	if uang <= 0 {
		return 0, errors.New("nilai setor anda tidak dikenali")
	}
	saldo = saldo + uang
	return saldo, nil
}

func (e EwalletImp) Print() float64 {
	return saldo
}

var saldoMock float64 = 0

type EwalletTest struct{}

func (e EwalletTest) Withdraw(uang float64) (float64, error) {
	if saldoMock < uang {
		return 0, errors.New("saldo anda tidak mencukupi")
	}
	saldoMock = saldoMock - uang
	return saldoMock, nil
}

func (e EwalletTest) Deposit(uang float64) (float64, error) {
	if uang <= 0 {
		return 0, errors.New("nilai setor anda tidak dikenali")
	}
	saldoMock = saldoMock + uang
	return saldoMock, nil
}

func (e EwalletTest) Print() float64 {
	return saldoMock
}

func JalankanAplikasiEwallet() Ewallet {
	if isTest {
		return EwalletTest{}
	}
	return EwalletImp{}
}

func Run(perintah []string) (float64, error) {
	var err error
	e := JalankanAplikasiEwallet()

	for _, val := range perintah {
		if val == "deposit" {
			_, err = e.Deposit(50000)
		} else {
			_, err = e.Withdraw(25000)
		}
	}
	return e.Print(), err
}

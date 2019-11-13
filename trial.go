package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"flag"
)

func main() {
	//fmt.Println("Hei, verden!")
	amount := flag.Int("amount", 100, "amount to pay")
	interestRate := flag.Float64("rate", 0.01, "interest rate to apply")
	flag.Parse()
	ps, err := pay(*amount, float32(*interestRate))
	if err != nil {
		log.Printf("%v", err)
		os.Exit(1)
		return
	}
	fmt.Printf("You paid a fee of %v (%T) on %v (%T) \n", ps.feeAmount, ps.feeAmount, ps.amount, ps.amount)
}

type paymentSummary struct{
	amount int
	feeAmount float32
}

func pay(amount int, fee float32) (paymentSummary, error) {
	if fee > 0.9 {
		return paymentSummary{}, errors.New("fee is too high")
	}
	feeAmount := float32(amount) * fee
	ps := paymentSummary{
		amount: amount,
		feeAmount: feeAmount,
	}
	return ps, nil
	//fmt.Println(amount)
}
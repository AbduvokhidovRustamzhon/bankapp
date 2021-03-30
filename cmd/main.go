package main

import (
	
	"bank/pkg/bank/payment"
	"bank/pkg/bank/types"
	"fmt"
)


func main() {
	// cards := []types.Card{
	// 	{
	// 		Balance: 1000,
	// 		Active: true,
	// 	},
	// 	{
	// 		Balance: 1000,
	// 		Active: false,
	// 	},
	// 	{
	// 		Balance: 1000,
	// 		Active: true,
	// 	},
	// }

	// // card.Deposit(&cards, 5_000_000)

	// sum := card.Total(cards)


	// fmt.Println(sum)
	paymentns := []types.Payment{
		{
			ID: 1,
			Amount: 100,
		},
		{
			ID: 2,
			Amount: 200,
		},
		{
			ID: 3,
			Amount: 500,
		},
		{
			ID: 4,
			Amount: 300,
		},

	}


	Newpayment := payment.Max(paymentns)
	fmt.Println(Newpayment)



}
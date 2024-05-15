package main

// MAIN IS THE ENTRY POINT FOR THIS MODULE - REQUIRED!!

// USE THE MAIN PACAKGE FOR ALL FILES IN THIS MODULE

import (
	"fmt"
)

func runProfitCalculator() {
	var revenue, expenses, taxRate float64
	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter Tax Rate: ")
	fmt.Scan(&taxRate)

	calculateProfit(revenue, expenses, taxRate)
}

func calculateProfit(revenue, expenses, taxRate float64) (float64, float64) {

	earningsBeforeTax := revenue - expenses
	taxAmount := earningsBeforeTax * taxRate
	earningsAfterTax := earningsBeforeTax - taxAmount

	fmt.Printf("Earnings Before Tax: %v \n", earningsBeforeTax)
	fmt.Printf("Tax Amount: %v \n", taxAmount)
	fmt.Printf("Earnings After Tax (Profit): %v\n", earningsAfterTax)

	return earningsBeforeTax, earningsAfterTax

}

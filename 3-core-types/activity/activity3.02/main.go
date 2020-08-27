package main

import (
	"errors"
	"fmt"
)

/*
	In this activity, we create a loan calculator for an online financial advisor platform.
	Our calculator has the following rules:
	1. A good credit score is a score of 450 above
	2. For a good credit score, your interest rate is 15%
	3. For a less than good score, your interest rate is 20%
	4. For a good credit score, your monthly payment must be no more than 20% of your monthly income
	5. For a less than good credit score, your monthly payment must be no more than 10% of your monthly income
	6. If a credit score, monthly income, loaan amount, or loan term is less than0, return an error
	7. If the term of the loan if not divisible by 12 months, return an error
	8. The interest payment will be a simple calculation of (loan amount * interest rate) / loan term
	9. After doing these calculations, display the following details to the user
*/

const (
	goodScore       = 450
	lowScoreRation  = 10
	goodScoreRation = 20
)

var (
	ErrCreditScore = errors.New("invalid credit score")
	ErrIncome      = errors.New("income invalid")
	ErrLoanAmount  = errors.New("loan amount invalid")
	ErrLoanTerm    = errors.New("loan term not a multiple of 12")
)

func checkLoan(creditScore int, income float64, loanAmount float64, loanTerm float64) error {
	interest := 20.0

	if creditScore >= goodScore {
		interest = 15.0
	}

	if creditScore < 1 {
		return ErrCreditScore
	}

	if income < 1 {
		return ErrIncome
	}

	if loanAmount < 1 {
		return ErrLoanAmount
	}

	if loanTerm < 1 || int(loanTerm)%12 != 0 {
		return ErrLoanTerm
	}

	rate := interest / 100

	payment := ((loanAmount * rate) / loanTerm) + (loanAmount / loanTerm)

	totalInterest := (payment * loanTerm) - loanAmount

	approved := false

	if income > payment {
		ratio := (payment / income) * 100

		if creditScore >= goodScore && ratio < goodScoreRation {
			approved = true
		} else if ratio < lowScoreRation {
			approved = true
		}
	}
	fmt.Println("Credit Score :", creditScore)
	fmt.Println("Income :", income)
	fmt.Println("Loan Amount :", loanAmount)
	fmt.Println("Loan Term :", loanTerm)
	fmt.Println("Monthly Payment :", payment)
	fmt.Println("Rate :", interest)
	fmt.Println("Total Cost :", totalInterest)
	fmt.Println("Approved :", approved)
	fmt.Println("")

	return nil
}

func main() {
	// Approved
	fmt.Println("Applicant 1")
	fmt.Println("-----------")
	err := checkLoan(500, 1000, 1000, 24)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Denied
	fmt.Println("Applicant 2")
	fmt.Println("-----------")
	err = checkLoan(350, 1000, 10000, 12)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

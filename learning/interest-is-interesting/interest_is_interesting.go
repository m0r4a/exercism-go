package interest

const (
	lessThan0          = 3.213
	between0and1000    = 0.5
	between1000and5000 = 1.621
	moreThan5000       = 2.475
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interest float32

	switch {
	case balance < 0:
		interest = lessThan0

	case 0 <= balance && balance < 1000:
		interest = between0and1000

	case 1000 <= balance && balance < 5000:
		interest = between1000and5000

	default:
		interest = moreThan5000
	}

	return interest
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	interest := InterestRate(balance)

	return balance * (float64(interest) / 100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	interest := float64(InterestRate(balance) / 100)

	return balance * (1 + interest)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var iterations int

	for balance < targetBalance {
		balance += (AnnualBalanceUpdate(balance) - balance)
		iterations += 1
	}

	return iterations
}

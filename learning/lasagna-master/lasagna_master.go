package lasagna

// PreparationTime recieves the layers of the lasagna and
// the average preparation time for each layer and calculates
// the preparation time for the lasagna.
func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2
	}

	return len(layers) * avgTime
}

// Quantities recieves a list of layers of a lasagna and
// returns the amout of noodles and sauce needed for it.
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64

	for _, value := range layers {
		if value == "noodles" {
			noodles += 50
		} else if value == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// AddSecretIngredient recieves your list of ingredients and
// you friend's list and returns your list of ingredients
// with you friend's secret ingredient added.
func AddSecretIngredient(friendsList, myList []string) {
	if len(friendsList) > 0 && len(myList) > 0 {
		secretIngredient := friendsList[len(friendsList)-1]
		myList[len(myList)-1] = secretIngredient
	}
}

// ScaleRecipe take two parameters, amounts and quantities and returns
// a slice of the amounts needed for the desired number of portions.
func ScaleRecipe(amounts []float64, quantities int) []float64 {
	var values []float64

	for _, v := range amounts {
		value := (float64(quantities) / 2) * v
		values = append(values, value)
	}

	return values
}

package lasagna

// PreparationTime get the total time
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return time * len(layers)
}

// Quantities get the noodles/sauce
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		} else if v == "sauce" {
			sauce += 0.2
		}
	}

	return
}

// AddSecretIngredient merge the list
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList) -1] = friendsList[len(friendsList) -1]
}

// ScaleRecipe scale the recipes
func ScaleRecipe(amounts []float64, num int) []float64 {
	s := make([]float64, len(amounts))

	for i := range amounts {
		s[i] = amounts[i] * float64(num) / 2.0
	}

	return s
}

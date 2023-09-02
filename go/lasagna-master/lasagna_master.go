package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(array []string, time int) (sum int) {
	if time == 0 {
		time = 2
	}
	sum = len(array) * time
	return
}

// TODO: define the 'Quantities()' function
func Quantities(slice []string) (grams int, sauce float64) {
	for i := 0; i < len(slice); i++ {
		if slice[i] == "noodles" {
			grams += 50
		} else if slice[i] == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(recipe []float64, portions int) (newRecipe []float64) {
	for i := 0; i < len(recipe); i++ {
		newRecipe = append(newRecipe, recipe[i]/2*float64(portions))
	}
	return
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

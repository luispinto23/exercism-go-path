// Package lasagna provides valuable methods to help to cook a brilliant lasagna.
package lasagna

// OvenTime returns how many minutes the lasagna should be in the oven.
func OvenTime() int {
	return 40
}

// RemainingOvenTime returns how many minutes the lasagna still has to remain in the oven, based on the expected oven time and the received time already spen in the oven
func RemainingOvenTime(timeInOven int) int {
	return OvenTime() - timeInOven
}

// PreparationTime returns how many minutes were spent preparing the lasagna, based on the inputed number of layers, assuming each layer takes 2 minutes to prepare.
func PreparationTime(layers int) int {
	return layers * 2
}

// ElapsedTime returns the overall time spent working on the lasagna, based on the number of layers and the number of minutes already passed in the oven.
func ElapsedTime(layers, minutesInOven int) int {
	return PreparationTime(layers) + minutesInOven
}

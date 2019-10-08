package main

func main() {
	keywords := []string{"chicken", "tacos", "avocado"}
	printMatchesForMeal("Today", "Breakfast", keywords)
	printMatchesForMeal("Today","Lunch", keywords)
	printMatchesForMeal("Today","Dinner", keywords)
}